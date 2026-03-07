# go-chinesecurrency [![Go version](https://img.shields.io/github/go-mod/go-version/tiger586/go-chinesecurrency)](https://github.com/tiger586/go-chinesecurrency/blob/main/go.mod)
將數字轉換為中文大寫財務貨幣格式（如「壹萬貳仟...」）

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
	fmt.Println(currency.ToChineseAmount(6666))
    // 輸出：陸仟陸佰陸拾陸
}
```

## 或可參考 example/main.go
```bash
go run example/main.go
```

## 測試
```bash
go test -v
```
