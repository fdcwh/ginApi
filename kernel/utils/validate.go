package utils

import (
	"github.com/go-playground/validator/v10"
	"goGIn/kernel"
	"reflect"
	"regexp"
)

// ValidatorGetError 自定义错误消息
func ValidatorGetError(ShouldError error, r interface{}) string {
	errors := ShouldError.(validator.ValidationErrors)
	if errors != nil {
		s := reflect.TypeOf(r)
		for _, fieldError := range errors {
			filed, _ := s.FieldByName(fieldError.Field())
			errTag := fieldError.Tag() + "_err"
			// 获取对应binding得错误消息
			errTagText := filed.Tag.Get(errTag)
			if errTagText != "" {
				return errTagText
			}
			// 获取统一错误消息
			errText := filed.Tag.Get("err")
			if errText != "" {
				return errText
			}
			// return fieldError.Field() + ":" + fieldError.Tag()
			return fieldError.Translate(kernel.FdTrans)
		}
	}
	return ""
}

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
