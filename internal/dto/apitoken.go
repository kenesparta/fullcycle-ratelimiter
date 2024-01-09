package dto

import "time"

type APIKeyRequestDB struct {
	MaxRequests   int     `json:"max_requests"`
	TimeWindowSec int64   `json:"time_window_sec"`
	Requests      []int64 `json:"requests"`
}

type APIKeyRequestSave struct {
	Value   string
	TimeAdd time.Time
}

type APIKeyInput struct {
	MaxRequests     int   `json:"max_requests"`
	TimeWindowSec   int64 `json:"time_window"`
	BlockedDuration int64 `json:"blocked_duration"`
}

type APIKeyCreateOutput struct {
	Key string `json:"api_key"`
}

type APIKeyOutput struct {
	Allow bool
}
