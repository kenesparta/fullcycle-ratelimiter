package entity

import "errors"

var (
	ErrIPExceededAmountRequest = errors.New("ip has exceeded the maximum amount of request allowed")
)
