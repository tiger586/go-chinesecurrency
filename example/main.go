package main

import (
	"fmt"

	currency "github.com/tiger586/go-chinesecurrency"
)

func main() {
	numbers := []int64{12345, 10005, 10500, 100000001, -100, 123456789}

	for _, n := range numbers {
		fmt.Printf("%d -> %s\n", n, currency.ToChineseAmount(n))
	}

	fmt.Println(currency.ToChineseAmount(6666))
}
