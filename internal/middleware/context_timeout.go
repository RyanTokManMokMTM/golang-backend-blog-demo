package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

//ContextTimeOut timeout control - router link A->B->C->D within time limit
func ContextTimeOut(t time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(ctx.Request.Context(), t) //set the request context with time out
		defer cancel()                                             //if time is out ,cancel the request

		ctx.Request = ctx.Request.WithContext(c) //change the context to contextWithTimeout of the request
		ctx.Next()
	}
}
