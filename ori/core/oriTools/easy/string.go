package easy

import "strings"

// 对手机号进行隐藏
func PrivateMobile(mobile string) string {
	masked := ""
	if len(mobile) > 5 {
		masked = mobile[:3] + strings.Repeat("*", len(mobile)-5) + mobile[len(mobile)-2:]
	}
	return masked
}
