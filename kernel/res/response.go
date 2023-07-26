package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (r *Res) ToJson(c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func (r *Res) ToXML(c *gin.Context) {
	c.XML(http.StatusOK, r)
}

// SetMsg 自定义响应信息
func (res *Res) SetMsg(message string) *Res {
	return &Res{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// SetData 追加响应数据
func (r *Res) SetData(data interface{}) *Res {
	return &Res{
		Code: r.Code,
		Msg:  r.Msg,
		Data: data,
	}
}

// SetCode  追加响应code
func (r *Res) SetCode(code int) *Res {
	return &Res{
		Code: code,
		Msg:  r.Msg,
		Data: r.Data,
	}
}

// ResVal 构造函数
func ResVal(code int, msg string) *Res {
	return &Res{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
