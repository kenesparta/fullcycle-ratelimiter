package entity

import "time"

// APIToken Thi token is related with the token API that we can generate for each client
type APIToken struct {
	Value     string
	Duration  time.Duration
	RateLimit int64
}

func (at *APIToken) DurationSeconds() time.Duration {
	return at.Duration * time.Second
}
