package i18n

import (
	"embed"
	"encoding/json"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

//go:embed locales/*
var fs embed.FS
var XLanguage = "X-Language"

func I18nHandler() gin.HandlerFunc {
	return ginI18n.Localize(
		ginI18n.WithBundle(&ginI18n.BundleCfg{
			DefaultLanguage:  language.Chinese,
			FormatBundleFile: "json",
			AcceptLanguage:   []language.Tag{language.English, language.Chinese},
			RootPath:         "locales/",
			UnmarshalFunc:    json.Unmarshal,
			// After commenting this line, use defaultLoader
			// it will be loaded from the file
			Loader: &ginI18n.EmbedLoader{
				FS: fs,
			}}),
		ginI18n.WithGetLngHandle(func(c *gin.Context, defaultLng string) (lang string) {
			// 从请求头中获取自定义的语言标识
			lang = c.Request.Header.Get(XLanguage)
			if lang == "" {
				// 如果请求头中没有，使用URL中的参数
				lang = c.Query("lang")
			}
			if lang == "" {
				// 也可以考虑默认从请求头中的 "Accept-Language" 获取
				accept := c.Request.Header.Get("Accept-Language")
				if accept != "" {
					tags, _, _ := language.ParseAcceptLanguage(accept)
					if len(tags) > 0 {
						lang = tags[0].String()
					}
				}
			}
			return
		}),
	)
}
