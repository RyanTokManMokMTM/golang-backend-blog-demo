package middleware

import (
	"errors"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		//token sent as query parameter or set inside the header?
		if s, exist := ctx.GetQuery("token"); exist {
			token = s
		} else {
			token = ctx.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			//parse the token
			_, err := app.ParseToken(token) //no need claims info in this case
			if err != nil {
				//cast to jwt error
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired: //ValidationError include lots of cases
					ecode = errcode.UnauthorizedTokenTimeOut
				default:
					ecode = errcode.UnauthorizedTokenError

				}
			}

		}

		//parse failed or token failed
		if !errors.Is(ecode, errcode.Success) {
			res := app.NewResponse(ctx)
			res.ToErrorResponse(ecode)
			ctx.Abort() //end to request or cancel the request
			return
		}

		ctx.Next()
	}
}
