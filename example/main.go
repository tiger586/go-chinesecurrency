package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	currency "github.com/tiger586/go-chinesecurrency"
)

func main() {
	CallClear() // 一執行就先清空畫面

	// float64
	var price float64
	price = 10050.58
	fmt.Println("數字：", price)

	// 情況 A：預設 (傳統台灣中文)
	fmt.Println("預設：", currency.ToString(price))
	// 輸出：壹萬零伍拾圓零捌分
	fmt.Println("無單位：", currency.ToString(price).Raw())
	// 輸出：壹萬零伍拾零捌分
	fmt.Println("--------------------------")

	// 情況 A-1：預設 (傳統台灣中文)，只修改部分設定
	// 1. 取得目前的預設配置（因為 DefaultTraditional 是全域變數）
	cfg := currency.DefaultTraditional
	// 2. 只修改 Symbol 欄位
	cfg.Symbol = "元"
	// 3. 將修改後的配置套用回全域設定
	currency.SetLangConfig(cfg)

	fmt.Println("預設：", currency.ToString(price), "-> 圓 改 元")
	// 輸出：壹萬零伍拾圓零捌分
	fmt.Println("無單位：", currency.ToString(price).Raw())
	// 輸出：壹萬零伍拾零捌分
	fmt.Println("--------------------------")

	// 情況 B：切換為簡體
	currency.SetLangConfig(currency.Simplified)
	fmt.Println("簡體：", currency.ToString(price))
	// 輸出：壹万零伍拾元零捌分
	fmt.Println("無單位：", currency.ToString(price).Raw())
	// 輸出：壹万零伍拾零捌分
	fmt.Println("--------------------------")

	// 情況 C：完全自定義 (例如 main 中傳入)
	currency.SetLangConfig(currency.LangConfig{
		Digits:   []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"},
		Units:    []string{"", "十", "百", "千"},
		BigUnits: []string{"", "萬", "億", "兆"},
		Negative: "負",
		Symbol:   "元",
		SubUnits: []string{"毛", "分"},
		Whole:    "整",
		Dot:      "點",
	})
	fmt.Println("口語：", currency.ToString(price))
	// 輸出：一万零五十五毛八分
	fmt.Println("無單位：", currency.ToString(price).Raw())
	// 輸出：一万零五十五毛八分
}

// CallClear 清空終端機畫面
func CallClear() {
	var cmd *exec.Cmd
	// 根據不同的作業系統，執行不同的清空指令
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
