package entity

type IP struct {
	Value           string
	BlockedDuration int64
	RateLimit       RateLimiter
}
