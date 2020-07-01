package util

import (
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		zh := zh.New()
		uni := ut.New(zh, en)
		trans, _ = uni.GetTranslator("zh")
		_ = zhtranslations.RegisterDefaultTranslations(v, trans)
	}
}

var trans ut.Translator

func WrapValidationError(err error) string {
	result := err.Error()
	if err, ok := err.(validator.ValidationErrors); ok {
		errs := err.Translate(trans)
		result = ""
		for _, v := range errs {
			result += v + ","
		}
	}
	return result
}
