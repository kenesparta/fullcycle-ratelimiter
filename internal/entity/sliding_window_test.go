package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOldRequests(t *testing.T) {
	startTime := time.Date(2024, time.January, 1, 12, 34, 56, 0, time.UTC)

	tests := []struct {
		name        string
		rl          RateLimiter
		expectedLen int
	}{
		{
			name: "No Requests to remove",
			rl: RateLimiter{
				Requests: []time.Time{
					startTime,
					startTime.Add(1 * time.Second),
				},
				TimeWindowSec: 1,
				MaxRequests:   10,
			},
			expectedLen: 2,
		},
		{
			name: "4 Requests to remove",
			rl: RateLimiter{
				Requests: []time.Time{
					startTime.Add(-4 * time.Second),
					startTime.Add(-3 * time.Second),
					startTime.Add(-2 * time.Second),
					startTime.Add(-1 * time.Second),
					startTime,
					startTime.Add(1 * time.Second),
				},
				TimeWindowSec: 1,
				MaxRequests:   10,
			},
			expectedLen: 2,
		},
		{
			name: "1 Requests to remove",
			rl: RateLimiter{
				Requests: []time.Time{
					startTime.Add(-1 * time.Second),
					startTime,
				},
				TimeWindowSec: 1,
				MaxRequests:   10,
			},
			expectedLen: 1,
		},
	}

	for i := 0; i < len(tests); i++ {
		t.Run(tests[i].name, func(t *testing.T) {
			tests[i].rl.removeOldRequests(startTime)
			assert.Len(t, tests[i].rl.Requests, tests[i].expectedLen)
		})
	}
}

func TestAllow(t *testing.T) {
	startTime := time.Date(2024, time.January, 1, 12, 34, 56, 0, time.UTC)

	tests := []struct {
		name          string
		rl            RateLimiter
		expectedAllow bool
	}{
		{
			name: "allow",
			rl: RateLimiter{
				Requests: []time.Time{
					startTime.Add(-90 * time.Millisecond),
					startTime.Add(-80 * time.Millisecond),
					startTime.Add(-70 * time.Millisecond),
					startTime.Add(-60 * time.Millisecond),
					startTime.Add(-50 * time.Millisecond),
					startTime.Add(-40 * time.Millisecond),
					startTime.Add(-30 * time.Millisecond),
					startTime.Add(-20 * time.Millisecond),
					startTime.Add(-10 * time.Millisecond),
					startTime,
				},
				TimeWindowSec: 1,
				MaxRequests:   10,
			},
			expectedAllow: true,
		},
		{
			name: "allow",
			rl: RateLimiter{
				Requests: []time.Time{
					startTime.Add(-14 * time.Second),
					startTime.Add(-13 * time.Second),
					startTime.Add(-12 * time.Second),
					startTime.Add(-11 * time.Second),
					startTime.Add(-10 * time.Second),
					startTime.Add(-9 * time.Second),
					startTime.Add(-8 * time.Second),
					startTime.Add(-7 * time.Second),
					startTime.Add(-6 * time.Second),
					startTime.Add(-5 * time.Second),
					startTime.Add(-4 * time.Second),
					startTime.Add(-3 * time.Second),
					startTime.Add(-2 * time.Second),
					startTime.Add(-1 * time.Second),
					startTime,
				},
				TimeWindowSec: 10,
				MaxRequests:   11,
			},
			expectedAllow: true,
		},
		{
			name: "no allow",
			rl: RateLimiter{
				Requests: []time.Time{
					startTime,
					startTime.Add(10 * time.Millisecond),
					startTime.Add(20 * time.Millisecond),
					startTime.Add(30 * time.Millisecond),
					startTime.Add(40 * time.Millisecond),
					startTime.Add(50 * time.Millisecond),
					startTime.Add(60 * time.Millisecond),
					startTime.Add(70 * time.Millisecond),
					startTime.Add(80 * time.Millisecond),
					startTime.Add(90 * time.Millisecond),
					startTime.Add(100 * time.Millisecond),
				},
				TimeWindowSec: 1,
				MaxRequests:   10,
			},
			expectedAllow: false,
		},
		{
			name: "no allow",
			rl: RateLimiter{
				Requests: func() []time.Time {
					var timeSlice []time.Time
					for i := 0; i < 100; i++ {
						timeSlice = append(timeSlice, startTime.Add(time.Duration(i)*time.Second))
					}
					return timeSlice
				}(),
				TimeWindowSec: 1,
				MaxRequests:   100,
			},
			expectedAllow: true,
		},
		{
			name: "allow",
			rl: RateLimiter{
				Requests: func() []time.Time {
					var timeSlice []time.Time
					for i := 0; i < 1000; i++ {
						timeSlice = append(timeSlice, startTime.Add(time.Duration(i)*time.Second))
					}
					return timeSlice
				}(),
				TimeWindowSec: 1,
				MaxRequests:   1000,
			},
			expectedAllow: true,
		},
	}

	for i := 0; i < len(tests); i++ {
		t.Run(tests[i].name, func(t *testing.T) {
			assert.Equal(t, tests[i].rl.Allow(startTime), tests[i].expectedAllow)
		})
	}
}
