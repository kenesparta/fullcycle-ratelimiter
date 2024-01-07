package dto

import "time"

type IPRequestSave struct {
	IP      string
	TimeAdd time.Time
}

type IPRequestResult struct {
	Allow bool
}

type IPRequestDB struct {
	MaxRequests   int     `json:"max_requests"`
	TimeWindowSec int64   `json:"time_window_sec"`
	Requests      []int64 `json:"requests"`
}
