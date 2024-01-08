package entity

import "errors"

var (
	ErrExceededRequest          = errors.New("you have reached the maximum number of Requests or actions allowed within a certain time frame")
	ErrIPExceededAmountRequest  = errors.New("ip has exceeded the maximum amount of request allowed")
	ErrAPIExceededAmountRequest = errors.New("this API key has exceeded the maximum amount of request allowed")

	ErrAPIKeyExceededAmountRequest = errors.New("your API key has exceeded the maximum amount of request allowed")

	ErrBlockedTimeDuration    = errors.New("blocked time duration should be greater than zero")
	ErrRateLimiterTimeWindow  = errors.New("rate limiter time window duration should be greater than zero")
	ErrRateLimiterMaxRequests = errors.New("rate limiter maximum requests should be greater than zero")
)
