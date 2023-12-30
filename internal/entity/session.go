package entity

import "time"

type RateLimiter struct {
	Requests    []time.Time
	TimeWindow  time.Duration
	MaxRequests int64
}

type Session struct {
	ID          string
	IP          string
	JWT         string
	APIToken    string
	RateLimiter RateLimiter
}
