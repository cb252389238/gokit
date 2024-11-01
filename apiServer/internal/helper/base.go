package helper

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

// EncodeUserPass 加密用户密码
func EncodeUserPass(password string) string {

	return fmt.Sprintf("%x", sha1.Sum([]byte(password)))
}

func GetUserId(c *gin.Context) string {
	return c.GetString("userId")
}

// 把斜杠转义符改成斜杠
func RemoveEscapedSlashes(s string) string {
	// 使用正则表达式匹配并替换掉转义斜杠
	re := regexp.MustCompile(`\\/`)
	return re.ReplaceAllString(s, "/")
}

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
		masked = idNo[0:4] + strings.Repeat("*", len(idNo)-6) + idNo[len(idNo)-2:]
	case len(idNo) > 6 && len(idNo) <= 14:
		masked = idNo[0:3] + strings.Repeat("*", len(idNo)-5) + idNo[len(idNo)-2:]
	case len(idNo) <= 6 && len(idNo) > 3:
		masked = strings.Repeat("*", len(idNo)-2) + idNo[len(idNo)-2:]
	default:
		masked = strings.Repeat("*", len(idNo))
	}
	return masked
}
