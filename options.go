package ratelimit

import "time"

type RateLimitOptions struct {
	limit time.Duration
}

type RateLimitOption func(*RateLimitOptions)

func Limit(t time.Duration) RateLimitOption {
	return func(o *RateLimitOptions) {
		o.limit = t
	}
}
