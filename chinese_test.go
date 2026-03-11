package currency

import (
	"testing"
)

func TestToChineseAmount(t *testing.T) {
	tests := []struct {
		name        string
		config      LangConfig
		input       float64
		expected    string
		expectedRaw string // 新增：預期去除單位後的樣子
	}{
		{
			name:        "基礎正體整數",
			config:      DefaultTraditional,
			input:       12345,
			expected:    "壹萬貳仟參佰肆拾伍圓整",
			expectedRaw: "壹萬貳仟參佰肆拾伍",
		},
		{
			name:        "正體小數(角分)",
			config:      DefaultTraditional,
			input:       123.4,
			expected:    "壹佰貳拾參圓肆角",
			expectedRaw: "壹佰貳拾參肆角",
		},
		{
			name:        "中間零且有分(補零邏輯)",
			config:      DefaultTraditional,
			input:       10.05,
			expected:    "壹拾圓零伍分",
			expectedRaw: "壹拾零伍分",
		},
		{
			name:        "10050測試",
			config:      DefaultTraditional,
			input:       10050.08,
			expected:    "壹萬零伍拾圓零捌分",
			expectedRaw: "壹萬零伍拾零捌分",
		},
		{
			name:        "跨萬位連續零",
			config:      DefaultTraditional,
			input:       10000500.1,
			expected:    "壹仟萬零伍佰圓壹角",
			expectedRaw: "壹仟萬零伍佰壹角",
		},
		{
			name:        "純小數",
			config:      DefaultTraditional,
			input:       0.58,
			expected:    "伍角捌分",
			expectedRaw: "伍角捌分", // 純小數本來就沒圓整
		},
		{
			name:        "負數測試",
			config:      DefaultTraditional,
			input:       -500.12,
			expected:    "負伍佰圓壹角貳分",
			expectedRaw: "負伍佰壹角貳分",
		},
		{
			name:        "簡體切換測試",
			config:      Simplified,
			input:       10005.6,
			expected:    "壹万零伍元陆角",
			expectedRaw: "壹万零伍陆角",
		},
		{
			name:        "零值測試",
			config:      DefaultTraditional,
			input:       0,
			expected:    "零圓整",
			expectedRaw: "零",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLangConfig(tt.config)

			actual := ToChineseAmount(tt.input)

			// 1. 測試預設財務格式 (含 圓/整)
			if string(actual) != tt.expected {
				t.Errorf("\n[財務格式錯誤] 輸入: %f\n期待: %s\n實際: %s", tt.input, tt.expected, string(actual))
			}

			// 2. 測試 Raw 格式 (不含 圓/整)
			if actual.Raw() != tt.expectedRaw {
				t.Errorf("\n[Raw格式錯誤] 輸入: %f\n期待: %s\n實際: %s", tt.input, tt.expectedRaw, actual.Raw())
			}
		})
	}
}

func TestConcurrentSafety(t *testing.T) {
	// 保持不變，檢查是否有 Data Race
	done := make(chan bool)
	go func() {
		for range 100 {
			SetLangConfig(Simplified)
			_ = ToChineseAmount(100.5).Raw()
		}
		done <- true
	}()
	go func() {
		for range 100 {
			SetLangConfig(DefaultTraditional)
			_ = ToChineseAmount(200).String()
		}
		done <- true
	}()
	<-done
	<-done
}
