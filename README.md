<h1 align="center">go-chinesecurrency</h1>
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/tiger586/go-chinesecurrency" />
<img src="https://img.shields.io/github/v/tag/tiger586/go-chinesecurrency?label=version"/>
<img src="https://img.shields.io/github/license/tiger586/go-chinesecurrency" />
</p>


這是一個輕量、精準且靈活的 Go 語言套件，專門用於將數字轉換為**中文大寫財務數字格式**（如支票、傳票、報表使用）。
支援**繁體/簡體**切換、自動處理**連續零**的讀法，並提供彈性的鏈式調用來決定是否顯示貨幣單位。
## 🌟 核心特色

* **精準補零**：完美處理 `10050` (壹萬零伍拾圓整) 與 `10.05` (壹拾圓零伍分) 等複雜補零邏輯。
* **正簡切換**：預設為繁體中文（財務大寫），一鍵即可切換為簡體中文。
* **彈性輸出**：預設輸出完整財務格式（圓、整），透過 `.Raw()` 即可取得純數字大寫。
* **線程安全**：內部使用 `sync.RWMutex` 確保在高併發環境下（如 Web Server）切換配置的安全。

## 📦 安裝方式

```bash
go get -u github.com/tiger586/go-chinesecurrency
```

## 🚀 快速上手
### 1. 基本用法 (預設傳統中文台灣財務格式)

```go
package main

import (
	"fmt"

	currency "github.com/tiger586/go-chinesecurrency"
)

func main() {
	num := 10050.08
	
	// 預設輸出：壹萬零伍拾圓零捌分
	fmt.Println(currency.ToChineseAmount(num))

	// 取得純數字大寫 (去除 圓、整)：壹萬零伍拾零捌分
	fmt.Println(currency.ToChineseAmount(num).Raw())
}
```

### 2. 切換為簡體中文

```go
// 切換全域設定為簡體
currency.SetLangConfig(currency.Simplified)

// 輸出：壹万零伍拾元零捌分
fmt.Println(currency.ToChineseAmount(10050.08))
// 取得純數字大寫 (去除 元、整)：壹万零伍拾零捌分
fmt.Println(currency.ToChineseAmount(10050.08).Raw())
```

### 3. 自定義語系或單位
> 如果您有特殊的顯示需求（例如口語化或特定幣別），可以自行定義 LangConfig：

```go
currency.SetLangConfig(currency.LangConfig{
	Digits:   []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"},
	Units:    []string{"", "十", "百", "千"},
	BigUnits: []string{"", "萬", "億", "兆"},
	Negative: "負",
	Symbol:   "元",
	SubUnits: []string{"毛", "分"},
	Whole:    "整",
})

// 輸出：一萬零五十元零八分
fmt.Println(currency.ToChineseAmount(10050.08))
// 取得純數字大寫 (去除 元、整)：一萬零五十零八分
fmt.Println(currency.ToChineseAmount(10050.08).Raw())
```

## 📝 範例參考 example/main.go
```bash
go run example/main.go
```

## 🧪 邏輯範例

| 輸入數字 | 預設輸出 (Financial) | .Raw() 輸出 |
| ---- | ---------------- | --------- |
| `12345` | 壹萬貳仟參佰肆拾伍圓整 | 壹萬貳仟參佰肆拾伍 |
| `10.05` | 壹拾圓零伍分 | 壹拾零伍分 |
| `10050.08` | 壹萬零伍拾圓零捌分 | 壹萬零伍拾零捌分 |
| `0.5` | 伍角 | 伍角 |
| `-123.4` | 負壹佰貳拾參圓肆角 | 負壹佰貳拾參肆角 |


## 🛠 運行測試
```bash
go test -v
```
