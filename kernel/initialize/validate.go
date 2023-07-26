package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"goGIn/kernel"
	"goGIn/kernel/utils"
	"log"
)

func InitValidator(locale string) {
	initTrans(locale)
	if v, res := binding.Validator.Engine().(*validator.Validate); res {
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)
	}
}

// initTrans 初始化翻译器
func initTrans(locale string) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		kernel.FdTrans, ok = uni.GetTranslator(locale)
		if !ok {
			log.Fatalf("uni.GetTranslator(%s) failed:", locale)
		}
		var err error
		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, kernel.FdTrans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, kernel.FdTrans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, kernel.FdTrans)
		}
		if err != nil {
			kernel.FdLog.Error("Init Trans failed, err:", zap.Error(err))
		}
	}
}
