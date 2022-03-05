package middleware

import (
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/RyanTokManMokMTM/blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
)

//RateLimiter pass any LimiterInterface(Limiter)
func RateLimiter(ml limiter.LimiterInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ml.Key(ctx) //get uri from context
		if b, ok := ml.GetBucket(key); ok {
			total := b.TakeAvailable(1) //return available token/removed bucket
			if total == 0 {
				//no Available token
				res := app.NewResponse(ctx)
				res.ToErrorResponse(errcode.TooManyRequest)
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
