package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"ori/internal/engine"
	"ori/internal/logic"
)

func Test1(oriEngine *engine.OriEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.Query("name") //127.0.0.1:9001/test1?name=111
		c.JSON(http.StatusOK, map[string]any{
			"code":  0,
			"msg":   "ok",
			"value": value,
		})
	}
}

func Test2(oriEngine *engine.OriEngine) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Test3(oriEngine *engine.OriEngine) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Test4(oriEngine *engine.OriEngine) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Test5(oriEngine *engine.OriEngine) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UserInfo(oriEngine *engine.OriEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Query("uid")
		if uid == "" {
			c.JSON(http.StatusOK, map[string]any{
				"code": 1001,
				"msg":  "参数错误",
			})
			return
		}
		fmt.Println(uid)
		user := new(logic.User)
		user.Uid = cast.ToInt(uid)
		info := user.GetUserInfo()
		c.JSON(http.StatusOK, map[string]any{
			"code":  0,
			"msg":   "ok",
			"value": info,
		})
		return
	}
}
