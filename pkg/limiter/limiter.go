//Package limiter - this middleware uses to limit the request amount
package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

//LimiterInterface define all function/property that limit is needed
type LimiterInterface interface {
	Key(ctx *gin.Context) string                    //get the bucket by key
	GetBucket(key string) (*ratelimit.Bucket, bool) //get rateLimit bucket
	AddBuckets(rule ...LimiterRule) LimiterInterface
}

//Limiter a group of rate limit bucket with a unique key
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterRule struct {
	Key          string        //key name
	FillInterval time.Duration // time to fill
	Capacity     int64         //bucket size
	Quantum      int64         //total number of token each interval
}
