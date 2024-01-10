package entity

import "errors"

var (
	ErrIPExceededAmountRequest     = errors.New("IP - you have reached the maximum number of Requests or actions allowed within a certain time frame - blocked")
	ErrAPIKeyExceededAmountRequest = errors.New("API TOKEN - you have reached the maximum number of Requests or actions allowed within a certain time frame - blocked")
	ErrBlockedTimeDuration         = errors.New("blocked time duration should be greater than zero")
	ErrRateLimiterTimeWindow       = errors.New("rate limiter time window duration should be greater than zero")
	ErrRateLimiterMaxRequests      = errors.New("rate limiter maximum requests should be greater than zero")
)
