package dto

import "time"

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
	KeyValue string `json:"api-key"`
}

type APIKeyOutput struct {
	Allow bool
}
