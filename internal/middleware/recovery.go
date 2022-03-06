package middleware

import (
	"fmt"
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/RyanTokManMokMTM/blog-service/pkg/mail"
	"github.com/gin-gonic/gin"
	"time"
)

//Recovery Custom Recover
//we also need to send out the email to notify developer
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			//what if panic happened?
			//recover
			//logging out the panic message

			if err := recover(); err != nil {
				s := "panic recover err: %v"
				//to know which function was panic
				global.Logger.WithCallerFrame().Errorf(s, err)

				//send emil
				m := mail.NewEmail(&mail.SMTP{
					Host:     global.EmailSetting.Host,
					Port:     global.EmailSetting.Port,
					IsSSL:    global.EmailSetting.IsSSL,
					UserName: global.EmailSetting.Email,
					Password: global.EmailSetting.Password,
					From:     global.EmailSetting.From,
				})

				err := m.SendMail(global.EmailSetting.To,
					fmt.Sprintf("Panic happend,occuring time:%d", time.Now().Unix()),
					fmt.Sprintf("Error Messgae %v", err))
				if err != nil {
					//email sending error
					global.Logger.Panicf("mail.SendMail error:%v", err)
				}

				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort() //end of the request
			}
		}()

		ctx.Next()
	}
}
