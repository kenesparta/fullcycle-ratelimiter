package dto

type APIKeyInput struct {
	MaxRequests     int   `json:"max_requests"`
	TimeWindowSec   int64 `json:"time_window"`
	BlockedDuration int64 `json:"blocked_duration"`
}

type APIKeyCreateOutput struct {
	TokenValue string `json:"api-key"`
}

type APIKeyOutput struct {
	Allow bool
}
