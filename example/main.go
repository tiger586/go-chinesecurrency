package main

import (
	"fmt"

	currency "github.com/tiger586/go-chinesecurrency"
)

func main() {
	// float64
	var price float64
	price = 6666

	// 情況 A：預設 (傳統台灣中文)
	fmt.Println("預設：", currency.ToChineseAmount(price))
	// 輸出：壹萬零伍拾圓零捌分
	fmt.Println("無單位：", currency.ToChineseAmount(price).Raw())
	// 輸出：壹萬零伍拾零捌分
	fmt.Println("--------------------------")

	// 情況 B：切換為簡體
	currency.SetLangConfig(currency.Simplified)
	fmt.Println("簡體：", currency.ToChineseAmount(price))
	// 輸出：壹万零伍拾元零捌分
	fmt.Println("無單位：", currency.ToChineseAmount(price).Raw())
	// 輸出：壹万零伍拾零捌分
	fmt.Println("--------------------------")

	// 情況 C：完全自定義 (例如 main 中傳入)
	currency.SetLangConfig(currency.LangConfig{
		Digits:   []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"},
		Units:    []string{"", "十", "百", "千"},
		BigUnits: []string{"", "万", "亿", "兆"},
		Negative: "负",
		Symbol:   "块",
		SubUnits: []string{"毛", "分"},
		Whole:    "整",
	})
	fmt.Println("口語：", currency.ToChineseAmount(price))
	// 輸出：一万零五十五毛八分
	fmt.Println("無單位：", currency.ToChineseAmount(price).Raw())
	// 輸出：一万零五十五毛八分
}
