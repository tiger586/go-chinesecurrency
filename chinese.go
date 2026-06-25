package currency

import (
	"math"
	"strings"
	"sync"
)

var (
	currentConfig = DefaultTraditional
	configMu      sync.RWMutex
)

type ChineseNumber string

// Raw 去除「圓」、「整」等單位，只保留純大寫數字
//
// 如果有小數時會加上 config 中設定的 Dot: "點" 或是 Dot: "点"，且小數點後數字不需要「角、分」
func (c ChineseNumber) Raw() string {
	configMu.RLock()
	cfg := currentConfig
	configMu.RUnlock()

	res := string(c)

	// 1. 如果包含「整」，代表是整數，直接去掉「圓」和「整」
	if strings.Contains(res, cfg.Whole) {
		res = strings.ReplaceAll(res, cfg.Symbol, "")
		res = strings.ReplaceAll(res, cfg.Whole, "")
		return res
	}

	// 2. 處理有小數的情況（動態檢查配置中的小數點單位，如「角」、「分」）
	hasSubUnit := false
	for _, subUnit := range cfg.SubUnits {
		if subUnit != "" && strings.Contains(res, subUnit) {
			hasSubUnit = true
			break
		}
	}

	if hasSubUnit {
		// 取得配置中的「點」，如果沒設定則預設用 "點"
		dotStr := cfg.Dot
		if dotStr == "" {
			dotStr = "點"
		}

		// 關鍵修正點：檢查字串裡有沒有「圓/元」
		if strings.Contains(res, cfg.Symbol) {
			// 有「圓」：正常將「圓 / 元」替換成配置的「點 / 点」
			res = strings.ReplaceAll(res, cfg.Symbol, dotStr)
		} else {
			// 沒有「圓」：代表是純小數（例如：伍角捌分），直接在最前面補上「零點」
			zeroStr := cfg.Digits[0]
			if zeroStr == "" {
				zeroStr = "零"
			}
			res = zeroStr + dotStr + res
		}

		// 循環拔除所有小數單位（例如：角、分）
		for _, subUnit := range cfg.SubUnits {
			if subUnit != "" {
				res = strings.ReplaceAll(res, subUnit, "")
			}
		}
		return res
	}

	// 3. 安全防線
	res = strings.ReplaceAll(res, cfg.Symbol, "")
	return res
}

// String 讓它直接支援 fmt.Println
func (c ChineseNumber) String() string {
	return string(c)
}

// SetLangConfig 允許全域切換語系設定
func SetLangConfig(config LangConfig) {
	configMu.Lock()
	defer configMu.Unlock()
	currentConfig = config
}

// ToChineseAmount 將數字轉換為中文財務大寫 (支援小數點兩位)
func ToString(f float64) ChineseNumber {
	// func ToChineseAmount(f float64) string {
	configMu.RLock()
	cfg := currentConfig
	configMu.RUnlock()

	if f == 0 {
		return ChineseNumber(cfg.Digits[0] + cfg.Symbol + cfg.Whole)
	}

	prefix := ""
	if f < 0 {
		prefix = cfg.Negative
		f = math.Abs(f)
	}

	// 四捨五入處理精度，轉為整數便於計算
	v := int64(math.Round(f * 100))
	intPart := v / 100
	decPart := v % 100

	var res []string

	// 1. 處理整數部分
	if intPart > 0 {
		res = append(res, formatInteger(intPart, cfg)+cfg.Symbol)
	}

	// 2. 處理小數部分
	if decPart > 0 {
		jiao := decPart / 10
		fen := decPart % 10

		if jiao > 0 {
			res = append(res, cfg.Digits[jiao]+cfg.SubUnits[0])
		} else if intPart > 0 {
			res = append(res, cfg.Digits[0]) // 10.05 -> 拾圓零伍分
		}

		if fen > 0 {
			res = append(res, cfg.Digits[fen]+cfg.SubUnits[1])
		}
	} else {
		res = append(res, cfg.Whole)
	}

	// return prefix + strings.Join(res, "")
	return ChineseNumber(prefix + strings.Join(res, ""))
}

func formatSection(n int64, c LangConfig) string {
	res := ""
	lastZero := true
	temp := n
	for i := range 4 {
		digit := temp % 10
		if digit == 0 {
			if !lastZero {
				res = c.Digits[0] + res
				lastZero = true
			}
		} else {
			res = c.Digits[digit] + c.Units[i] + res
			lastZero = false
		}
		temp /= 10
	}
	return res
}

func formatInteger(num int64, cfg LangConfig) string {
	if num == 0 {
		return cfg.Digits[0]
	}

	var sections []string
	unitIdx := 0
	tempNum := num
	needZero := false // 用來標記是否需要在下一節前面補「零」

	for tempNum > 0 {
		section := tempNum % 10000
		if section > 0 {
			s := formatSection(section, cfg)

			// 關鍵修正：如果這節不滿 1000 (如 50)，且這節不是最高位，
			// 或者前一節(低位)有數字且這節與低位之間有空位，就補「零」
			if needZero && !strings.HasPrefix(s, cfg.Digits[0]) {
				s = cfg.Digits[0] + s
			}

			sections = append([]string{s + cfg.BigUnits[unitIdx]}, sections...)

			// 如果這節「結尾」是零，或者這節不滿千位，下一節(更高位)若有數字則需補零
			needZero = (section < 1000)
		} else {
			// 這一節全為零 (如 0000 萬)，標記下一節需要補零
			if len(sections) > 0 {
				needZero = true
			}
		}
		tempNum /= 10000
		unitIdx++
	}

	result := strings.Join(sections, "")
	// 最終檢查：確保不會有連續兩個「零」
	return sanitize(result, cfg)
}

func sanitize(s string, c LangConfig) string {
	doubleZero := c.Digits[0] + c.Digits[0]
	for strings.Contains(s, doubleZero) {
		s = strings.ReplaceAll(s, doubleZero, c.Digits[0])
	}
	s = strings.TrimPrefix(s, c.Digits[0])
	if s == "" {
		return c.Digits[0]
	}
	return s
}

// ToChineseAmount 保留給舊用戶，請改用 ToString()
//
// Deprecated: Use ToString() instead.
func ToChineseAmount(amount float64) ChineseNumber {
	return ToString(amount)
}
