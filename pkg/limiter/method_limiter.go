package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

//MethodLimiter include a limiter and limiter manger rateLimit bucket 
type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterInterface {
	return MethodLimiter{&Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}}
}

//Key getting the key/token(uri path) of buckets
func (ml MethodLimiter) Key(ctx *gin.Context) string {
	//get the key from uri
	uri := ctx.Request.RequestURI
	i := strings.Index(uri, "?") //where is the query string index
	if i == -1 {
		//no query parameter
		return ""
	}

	return uri[:i] //return whole path except query string

}

func (ml MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := ml.limiterBuckets[key] //get the bucket by key/token
	return bucket, ok
}

func (ml MethodLimiter) AddBuckets(rule ...LimiterRule) LimiterInterface {
	//add to the bucket
	for _, bucketRule := range rule {
		//if not exist,append to the bucket list
		if _, ok := ml.limiterBuckets[bucketRule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(bucketRule.FillInterval, bucketRule.Capacity, bucketRule.Quantum)
			ml.limiterBuckets[bucketRule.Key] = bucket
		}
	}

	return ml
}
