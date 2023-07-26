package res

import (
	"net/http"
)

var (
	Success = ResVal(http.StatusOK, "success")                // 通用成功
	Error   = ResVal(http.StatusInternalServerError, "error") // 通用错误

	ErrParam = ResVal(422, "参数有误")

	ErrTokenParam   = ResVal(10001, "Token不能为空")
	ErrTokenInvalid = ResVal(10002, "Invalid token")
	ErrTokenExpired = ResVal(10003, "Expired token")
)
