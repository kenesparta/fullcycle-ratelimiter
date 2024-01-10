package entity

const (
	IPPrefixRateKey            = "rate:ip"
	IPPrefixBlockedDurationKey = "blocked:ip"
	StatusIPBlocked            = "IPBlocked"
)

type IP struct {
	value string

	// BlockedDuration is the number of SECONDS that it blocks the IP if it reaches the RateLimiter.MaxRequests each
	// RateLimiter.TimeWindowSec amount of seconds.
	BlockedDuration int64
	RateLimiter     RateLimiter
}

func (ip *IP) SaveValue(ipValue string) {
	ip.value = ipValue
}

func (ip *IP) Value() string {
	return ip.value
}
