package dto

import "time"

type RequestSave struct {
	IP       string
	APIToken string
	TimeAdd  time.Time
}

type RequestResult struct {
	Allow bool
}
