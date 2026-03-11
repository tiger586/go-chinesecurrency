# go-chinesecurrency [![Go version](https://img.shields.io/github/go-mod/go-version/tiger586/go-chinesecurrency)](https://github.com/tiger586/go-chinesecurrency/blob/main/go.mod)
將數字轉換為中文大寫財務貨幣格式（如「壹萬貳仟...」）。
- 可使用到小數兩位，
- 預設使用傳統中文台灣格式，
- 也可以切換為簡體中國模式，或是自訂文字內容。
- 預設顯示單位，也可只顯示數字
> 例如 105，預設：壹佰零伍圓整，無單位：壹佰零伍

## 安裝

```bash
go get -u github.com/tiger586/go-chinesecurrency
```

## 使用方法

```go
import (
	"fmt"

	currency "github.com/tiger586/go-chinesecurrency"
)

func main() {
	// 預設
	fmt.Println(currency.ToChineseAmount(6666))
    // 輸出：陸仟陸佰陸拾陸圓整
	
	// 無單位
	fmt.Println(currency.ToChineseAmount(6666).Raw())
    // 輸出：陸仟陸佰陸拾陸圓整
}
```

## 更多請參考 example/main.go
```bash
go run example/main.go
```

## 測試
```bash
go test -v
```
