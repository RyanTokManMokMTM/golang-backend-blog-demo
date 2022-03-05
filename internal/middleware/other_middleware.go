package middleware

import (
	"bytes"
	"fmt"
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/RyanTokManMokMTM/blog-service/pkg/logger"
	"github.com/RyanTokManMokMTM/blog-service/pkg/mail"
	"github.com/gin-gonic/gin"
	"time"
)

//AccessLogWriter Used to access response header
//Used to capture info with request ,response ,starting time, end time
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""), //storing info from here
			ResponseWriter: ctx.Writer,
		}

		ctx.Writer = bodyWriter
		start := time.Now().Unix()
		ctx.Next() //wait for it
		end := time.Now().Unix()

		field := logger.Fields{
			"request":  ctx.Request.PostForm.Encode(), //encoding the form data
			"response": bodyWriter.body.String(),
		}
		s := "access log: method:%s,statusCode:%d,begin_time:%d,end_time=%d"

		global.Logger.WithFields(field).Infof(s, ctx.Request.Method, bodyWriter.Status(), start, end)
	}
}

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

//AppInfo example of Service Info Storing for inner server
func AppInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//MetaData
		ctx.Set("Server version", "v0.0.1") //string:interface
		ctx.Set("Server name", "Blog-Service")
		ctx.Next()
	}
}
