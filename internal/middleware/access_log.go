package middleware

import (
	"bytes"
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/pkg/logger"
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
