package dto

type APITokenInput struct {
	MaxRequests     int   `json:"max_requests"`
	TimeWindowSec   int64 `json:"time_window"`
	BlockedDuration int64 `json:"blocked_duration"`
}

type APITokenOutput struct {
}
