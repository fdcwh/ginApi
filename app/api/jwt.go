package api

import (
	"github.com/gin-gonic/gin"
	"goGIn/kernel/res"
	"goGIn/kernel/utils"
	"strings"
)

func JwtSet(c *gin.Context) {
	var sys utils.JwtSysClaims
	sys.UserID = 1

	JwtTokenOut, err := utils.JwtGenToken(sys)
	if err != nil {
		res.Error.SetMsg(err.Error()).ToJson(c)
		return
	}

	res.Success.SetData(JwtTokenOut).ToJson(c)
}

func JwtGet(c *gin.Context) {
	token := c.GetHeader("Authorization")
	parts := strings.SplitN(token, " ", 2)

	JwtTokenOut, err := utils.JwtParseToken(parts[1])
	if err != nil {
		res.Error.SetMsg(err.Error()).ToJson(c)
		return
	}
	res.Success.SetData(gin.H{
		"u":           c.GetInt64("uid"),
		"JwtTokenOut": JwtTokenOut,
	}).ToJson(c)
}

func RefreshToken(c *gin.Context) {

	res.Success.ToJson(c)
}
