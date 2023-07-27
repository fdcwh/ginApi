package middleware

import (
	"github.com/gin-gonic/gin"
	"goGIn/kernel/res"
	"goGIn/kernel/utils"
	"strings"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取到请求头中的token
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			res.ErrTokenParam.ToJson(c)
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// 访问失败,无效的token,请登录!
			res.ErrTokenParam.ToJson(c)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := utils.JwtParseToken(parts[1])

		if err != nil {
			// 访问失败,无效的token,请登录!
			res.ErrTokenParam.ToJson(c)
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("uid", mc.JwtSysClaims.UserID)
		c.Next()
	}
}
