package exception

import "net/http"

const (
	UNKNOWN_ERROR = 1002 // 未知错误
)

// APIException api错误的结构体
type APIException struct {
	code int    `json:"error_code"`
	msg  string `json:"msg"`
	data string `json:"data"`
}

// 实现接口
func (e *APIException) Error() string {
	return e.msg
}

func newAPIException(errorCode int, msg string) *APIException {
	return &APIException{
		code: errorCode,
		msg:  msg,
	}
}

// ServerError 500 错误处理
func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// NotFound 404 错误
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

// UnknownError 未知错误
func UnknownError(message string) *APIException {
	return newAPIException(UNKNOWN_ERROR, message)
}
