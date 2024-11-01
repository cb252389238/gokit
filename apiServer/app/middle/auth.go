package middle

import (
	"apiServer/core/coreConfig"
	"apiServer/core/coreJwtToken"
	error2 "apiServer/i18n/error"
	"apiServer/typedef/response"
	"github.com/gin-gonic/gin"
)

// Auth
//
//	@Description:	权限认证
//	@return			gin.HandlerFunc
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO 进行token验证
		token := c.GetHeader("Authorization")
		claims, err := coreJwtToken.Decode(token, []byte(coreConfig.GetHotConf().JwtSecret))
		if err != nil {
			response.FailResponse(c, error2.I18nError{
				Code: error2.ErrorCodeToken,
				Msg:  nil,
			})
			c.Abort()
			return
		}
		userId := claims.UserId
		c.Set("userId", userId)
		c.Next()
	}
}
