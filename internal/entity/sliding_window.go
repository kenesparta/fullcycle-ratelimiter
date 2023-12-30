package entity

import (
	"sync"
	"time"
)

// RateLimiter is a struct that implements rate limiting logic.
// It's designed to limit the number of requests allowed within a specified time window.
type RateLimiter struct {
	// Requests is a slice of time.Time that holds the timestamps of the incoming requests.
	// It is used to keep track of the requests that have been made and whether a new request
	// should be allowed or not based on the timing of previous requests.
	Requests []time.Time

	// TimeWindowSec specifies the duration in SECONDS of the time window for which the requests are counted.
	// For instance, if TimeWindow is set to 60, the RateLimiter will only consider
	// the number of requests in the last minute.
	TimeWindowSec int64

	// MaxRequests defines the maximum number of requests that are allowed within the TimeWindow.
	// For example, we have TimeWindowSec = 1 and MaxRequests = 100, we have the max request limit 100 req/s
	// Other example, we have TimeWindowSec = 60 and MaxRequests = 100, we have the max request limit 100 req/min
	MaxRequests int

	// lock is a mutex that ensures that access to the Requests slice is synchronized across multiple goroutines.
	lock sync.Mutex
}

// GetDurationTimeWindow  converts the time window from an integer number of seconds into a time.Duration type.
func (rl *RateLimiter) GetDurationTimeWindow() time.Duration {
	return time.Duration(rl.TimeWindowSec) * time.Second
}

// RemoveOldRequests is a method of the RateLimiter struct that removes requests that are older than the current
// time window. This method assumes that rl.Requests is sorted in ascending order of request time.
func (rl *RateLimiter) RemoveOldRequests(fromTime time.Time) {
	threshold := fromTime.Add(-rl.GetDurationTimeWindow())
	start := 0
	for i, t := range rl.Requests {
		if t.After(threshold) {
			start = i
			break
		}
	}
	rl.Requests = rl.Requests[start:]
}

// Allow determines whether a new request at the given time should be allowed based on the rate limit policy.
func (rl *RateLimiter) Allow(fromTime time.Time) bool {
	rl.lock.Lock()
	defer rl.lock.Unlock()

	rl.RemoveOldRequests(fromTime)
	if len(rl.Requests) < rl.MaxRequests {
		rl.AddRequests(fromTime)
		return true
	}

	return false
}

func (rl *RateLimiter) AddRequests(request time.Time) {
	rl.Requests = append(rl.Requests, request)
}
