package middleware

import "github.com/gin-gonic/gin"

//AppInfo example of Service Info Storing for inner server
func AppInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//MetaData
		ctx.Set("Server version", "v0.0.1") //string:interface
		ctx.Set("Server name", "Blog-Service")
		ctx.Next()
	}
}
