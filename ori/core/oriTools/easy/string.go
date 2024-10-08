package easy

import "strings"

// 脱敏手机号
func PrivateMobile(mobile string) string {
	masked := ""
	if len(mobile) > 5 {
		masked = mobile[:3] + strings.Repeat("*", len(mobile)-5) + mobile[len(mobile)-2:]
	}
	return masked
}

// 脱敏名字
func PrivateRealName(name string) string {
	masked := ""
	runeName := []rune(name)
	switch {
	case len(runeName) > 3:
		masked = string(runeName[0]) + strings.Repeat("*", len(runeName)-2) + string(runeName[len(runeName)-1:])
	case len(runeName) == 3:
		masked = string(runeName[0]) + strings.Repeat("*", 1) + string(runeName[len(runeName)-1:])
	case len(runeName) == 2:
		masked = string(runeName[0]) + strings.Repeat("*", 1)
	default:
		masked = name
	}
	return masked
}

// 脱敏身份证号
func PrivateIdNo(idNo string) string {
	masked := ""
	switch {
	case len(idNo) > 14:
		masked = idNo[0:4] + strings.Repeat("*", len(idNo)-2) + idNo[len(idNo)-2:]
	case len(idNo) < 6:
		masked = idNo[0:2] + strings.Repeat("*", len(idNo)-2) + idNo[len(idNo)-2:]
	default:
		masked = idNo[0:4] + strings.Repeat("*", len(idNo)-2) + idNo[len(idNo)-2:]
	}
	return masked
}
