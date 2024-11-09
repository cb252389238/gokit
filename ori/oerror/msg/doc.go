package i18n_msg

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// 获取多语言返回
func GetI18nMsg(c *gin.Context, msgKey string, templateData ...map[string]any) string {
	var data interface{}
	if len(templateData) > 0 {
		data = templateData[0]
	}
	message := ginI18n.MustGetMessage(c, &i18n.LocalizeConfig{
		MessageID:    msgKey,
		TemplateData: data,
	})
	return message
}

type JoinMsg struct {
	MsgKey       string
	TemplateData map[string]any
}

// 拼接多条消息
func SprintfI18nMsg(c *gin.Context, data []JoinMsg) (msg string) {
	if len(data) == 0 {
		return
	}
	for _, v := range data {
		message := ginI18n.MustGetMessage(c, &i18n.LocalizeConfig{
			MessageID:    v.MsgKey,
			TemplateData: v.TemplateData,
		})
		msg += message
	}
	return
}
