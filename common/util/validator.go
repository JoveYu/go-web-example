package util

import (
	"errors"
	"strings"

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
		uni := ut.New(en, zh)
		trans, _ = uni.GetTranslator("zh")
		_ = zhtranslations.RegisterDefaultTranslations(v, trans)
	}
}

var trans ut.Translator

func WrapValidationError(err error) string {
	var result string
	verr := validator.ValidationErrors{}
	if errors.As(err, &verr) {
		errs := verr.Translate(trans)
		strs := []string{}
		for _, v := range errs {
			strs = append(strs, v)
		}
		result = strings.Join(strs, ",")
	} else {
		result = err.Error()
	}
	return result
}
