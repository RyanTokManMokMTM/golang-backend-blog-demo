//Package middleware - translate validator message
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

func Translation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//translator
		uni := ut.New(en.New(), zh.New(), zh_Hans.New())          //using en ,zh,zh_han translator
		local := ctx.GetHeader("local")                           //get location from header
		trans, _ := uni.GetTranslator(local)                      //get translator by local
		v, ok := binding.Validator.Engine().(*validator.Validate) //getting the binding validator
		if ok {
			switch local { //depend on local
			case "zh":
				_ = zh_trans.RegisterDefaultTranslations(v, trans) //register the translator by zh
			case "en":
				_ = en_trans.RegisterDefaultTranslations(v, trans) //register the translator by eng
			default:
				_ = zh_trans.RegisterDefaultTranslations(v, trans) //default zh
			}
			ctx.Set("trans", trans)
		}
		ctx.Next()
	}
}
