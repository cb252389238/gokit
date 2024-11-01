package easy

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

// DefaultLength 默认保留小数位
const DefaultLength int32 = 2

// StringToDecimal
//
//	@Description: 字符串转十进制小数
//	@param s string - 待转换的字符串
//	@return decimal.Decimal - 已转换的十进制小数
func StringToDecimal(s string) decimal.Decimal {
	d, _ := decimal.NewFromString(s)
	return d
}

// StringToDecimalFixed
//
//	@Description: 字符串保留几位小数(默认保留两位小数)
//	@param s string - 待格式化的字符串
//	@param length ...int32 - 保留几位小数
//	@return string - 已格式化小数位的字符串
func StringToDecimalFixed(s string, length ...int32) string {
	d, _ := decimal.NewFromString(s)
	defaultLen := DefaultLength
	if len(length) > 0 {
		defaultLen = length[0]
	}
	return d.StringFixed(defaultLen)
}

// StringFixed
//
//	@Description: 十进制小数转字符串类型(默认保留两位小数)
//	@param d decimal.Decimal - 格式化为字符串的十进制小数
//	@param length ...int32 - 保留几位小数
//	@return string - 十进制小数的字符串类型
func StringFixed(d decimal.Decimal, length ...int32) string {
	defaultLen := DefaultLength
	if len(length) > 0 {
		defaultLen = length[0]
	}
	return d.StringFixed(defaultLen)
}

// 格式化榜单值为 "W" 单位，并保留一位小数
func FormatLeaderboardValue(value float64) string {
	if value < 1 {
		return "0"
	}
	if value < 10000 {
		// 如果值小于 10000，直接返回原值（可以根据需求调整）
		return fmt.Sprintf("%.0f", value)
	}

	// 计算以 W 为单位的值
	wValue := value / 10000

	// 四舍五入保留一位小数
	roundedValue := math.Round(wValue*10) / 10

	// 返回格式化的字符串
	return fmt.Sprintf("%.1fW", roundedValue)
}
