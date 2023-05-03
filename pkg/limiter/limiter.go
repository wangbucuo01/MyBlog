package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	// 获取对应的限流器的键值对名称
	Key(c *gin.Context) string
	// 获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	// 新增多个令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

// 令牌桶与键值对名称的映射关系
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// 存储令牌桶的一些相应规则属性
type LimiterBucketRule struct {
	// 键值对名称
	Key          string
	// 间隔多久时间放N个令牌
	FillInterval time.Duration
	// 令牌桶的容量
	Capacity     int64
	// 每次到达间隔时间后所放的具体令牌数量
	Quantum      int64
}
