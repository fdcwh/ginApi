package exception

import (
	"github.com/gin-gonic/gin"
	"goGIn/kernel/res"
)

type HandlerFunc func(c *gin.Context) error

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)
		err = handler(c)
		if err != nil {
			var apiException *APIException
			if h, ok := err.(*APIException); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == "debug" {
					// 错误
					apiException = UnknownError(e.Error())
				} else {
					// 未知错误
					apiException = UnknownError(e.Error())
				}
			} else {
				apiException = ServerError()
			}
			apiException.msg = c.Request.Method + " " + c.Request.URL.String()
			res.Error.SetMsg(apiException.msg).ToJson(c)
			c.Abort()
			return
		}
	}
}

func HandleNotFound(c *gin.Context) {
	handleErr := NotFound()
	// handleErr.msg = c.Request.Method + " ---》》》" + c.Request.URL.String()
	res.Error.SetMsg(handleErr.msg + "【" + c.Request.Method + " -- " + c.Request.URL.String() + "】").SetCode(handleErr.code).ToJson(c)
	c.Abort()
	return
}
