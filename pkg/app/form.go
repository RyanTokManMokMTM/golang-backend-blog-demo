package app

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

//Error implement error interface and return message
func (v *ValidError) Error() string {
	return v.Message
}

//Error implement error interface and return combined message
func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

//Errors append all ValidErrors as error string to list
func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(ctx *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := ctx.ShouldBind(&v) //binding the context to an interface
	if err != nil {
		v := ctx.Value("trans")
		global.Logger.Warning("testing")
		trans, _ := v.(ut.Translator)
		vals, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs //no error but failed
		}

		for key, value := range vals.Translate(trans) { //translate err by trans(translator from ctx )
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs //with error
	}
	return true, nil //no error
}
