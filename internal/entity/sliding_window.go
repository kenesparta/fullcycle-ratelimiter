package entity

import (
	"sync"
	"time"
)

type SlidingWindow struct {
	// Requests Slice to store timestamps of requests
	Requests []time.Time

	// MaxRequest Max number of requests allowed in the window
	MaxRequest int

	// Window Duration of the sliding window
	window time.Duration
	lock   sync.Mutex
}

func (l *SlidingWindow) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now()
	// Remove old requests outside the window
	threshold := now.Add(-l.window)
	start := 0
	for i, t := range l.Requests {
		if t.After(threshold) {
			start = i
			break
		}
	}
	l.Requests = l.Requests[start:]

	if len(l.Requests) < l.MaxRequest {
		l.Requests = append(l.Requests, now)
		return true
	}
	return false
}
