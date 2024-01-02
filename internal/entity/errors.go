package entity

import "errors"

var (
	ErrExceededRequest          = errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")
	ErrIPExceededAmountRequest  = errors.New("ip has exceeded the maximum amount of request allowed")
	ErrAPIExceededAmountRequest = errors.New("this API token has exceeded the maximum amount of request allowed")
)
