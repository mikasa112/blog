package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := ctx.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v := binding.Validator.Engine().(*validator.Validate)
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		default:
			zh_translations.RegisterDefaultTranslations(v, trans)
		}
		ctx.Set("trans", trans)
		ctx.Next()
	}
}
