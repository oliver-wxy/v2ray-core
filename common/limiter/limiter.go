package limiter

import (
	"golang.org/x/time/rate"
	"sync"
)

type Limiter struct {
	Limit *sync.Map // Key: mail, Value: *InboundInfo
}

func New() *Limiter {
	return &Limiter{
		Limit: new(sync.Map),
	}
}

func (l *Limiter) GetUserBucket(email string, limite int32) (limiter *rate.Limiter, SpeedLimit bool) {
	// Speed limit
	if limite > 0 {
		limiter := rate.NewLimiter(rate.Limit(limite), int(limite)) // Byte/s
		return limiter, true
	} else {
		return nil, false
	}
}
