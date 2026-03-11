package currency

// LangConfig 定義中文大寫所需的字元與單位
type LangConfig struct {
	Digits   []string // 零-玖
	Units    []string // 拾、佰、仟
	BigUnits []string // 萬、億、兆
	Negative string   // 負
	Symbol   string   // 圓 / 元
	SubUnits []string // 角、分
	Whole    string   // 整
}

var (
	// DefaultTraditional 預設傳統中文台灣 (財務大寫)
	DefaultTraditional = LangConfig{
		Digits:   []string{"零", "壹", "貳", "參", "肆", "伍", "陸", "柒", "捌", "玖"},
		Units:    []string{"", "拾", "佰", "仟"},
		BigUnits: []string{"", "萬", "億", "兆"},
		Negative: "負",
		Symbol:   "圓",
		SubUnits: []string{"角", "分"},
		Whole:    "整",
	}

	// Simplified 簡體中文配置
	Simplified = LangConfig{
		Digits:   []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"},
		Units:    []string{"", "拾", "佰", "仟"},
		BigUnits: []string{"", "万", "亿", "兆"},
		Negative: "负",
		Symbol:   "元",
		SubUnits: []string{"角", "分"},
		Whole:    "整",
	}
)
