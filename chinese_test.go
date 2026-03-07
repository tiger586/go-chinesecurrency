package currency

import (
	"testing"
)

func TestToChineseAmount(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "零"},
		{5, "伍"},
		{10, "壹拾"},
		{100, "壹佰"},
		{1000, "壹仟"},
		{10000, "壹萬"},
		{12345, "壹萬貳仟參佰肆拾伍"},
		{10005, "壹萬零伍"},     // 測試中間連續零
		{10500, "壹萬零伍佰"},    // 測試末尾零
		{100005, "壹拾萬零伍"},   // 測試跨萬位零
		{100000001, "壹億零壹"}, // 測試跨億位零
		{102030405, "壹億零貳佰零參萬零肆佰零伍"},
		{-123, "負壹佰貳拾參"}, // 測試負數
		{123456789, "壹億貳仟參佰肆拾伍萬陸仟柒佰捌拾玖"},
	}

	for _, tt := range tests {
		t.Run(string(tt.expected), func(t *testing.T) {
			actual := ToChineseAmount(tt.input)
			if actual != tt.expected {
				t.Errorf("ToChineseAmount(%d) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}
