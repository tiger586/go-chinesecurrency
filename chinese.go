package currency

import (
	"strings"
)

var (
	zhDigits   = []string{"零", "壹", "貳", "參", "肆", "伍", "陸", "柒", "捌", "玖"}
	zhUnits    = []string{"", "拾", "佰", "仟"}
	zhBigUnits = []string{"", "萬", "億", "兆"}
)

func ToChineseAmount(num int64) string {
	if num == 0 {
		return zhDigits[0]
	}

	// 處理負數
	prefix := ""
	if num < 0 {
		prefix = "負"
		num = -num
	}

	var res []string
	unitCount := 0

	for num > 0 {
		section := num % 10000 // 每次處理四位數 (一小節)
		if section > 0 {
			s := formatSection(section)
			res = append([]string{s + zhBigUnits[unitCount]}, res...)
		} else if len(res) > 0 && !strings.HasPrefix(res[0], zhDigits[0]) {
			// 如果這一節全為零，且後面還有數字，需判斷是否補個「零」來連接
			res = append([]string{zhDigits[0]}, res...)
		}
		num /= 10000
		unitCount++
	}

	// 整理字串：處理開頭可能的零，以及多餘的連續零
	result := strings.Join(res, "")
	return prefix + strings.TrimPrefix(result, zhDigits[0])
}

func formatSection(n int64) string {
	s := ""
	temp := n
	lastZero := true // 預設為 true，避免末尾出現零

	for i := range 4 {
		digit := temp % 10
		if digit == 0 {
			if !lastZero { // 只有在之前有非零數字時，才補一個「零」
				s = zhDigits[0] + s
				lastZero = true
			}
		} else {
			s = zhDigits[digit] + zhUnits[i] + s
			lastZero = false
		}
		temp /= 10
	}
	return s
}
