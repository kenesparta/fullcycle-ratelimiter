package entity

import (
	"context"
)

//go:generate mockgen -source repository.go -destination mock/repository_mock.go -package mock
type commonRepo interface {
	// UpsertRequest Updates or inserts a new request inside the RateLimiter.requests array.
	// This also creates a new instance of Request
	UpsertRequest(ctx context.Context, key string, rl *RateLimiter) error

	// SaveBlockedDuration Stores the blocked duration amount by key
	SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error

	// GetBlockedDuration Obtain the blocked duration by key
	GetBlockedDuration(ctx context.Context, key string) (string, error)

	// GetRequest reads the stored array of request
	GetRequest(ctx context.Context, value string) (*RateLimiter, error)
}

type APITokenRepository interface {
	// Save Persists a new APIToken with the initial values of MaxRequest and TimeWindowSec
	Save(ctx context.Context, token *APIToken) error

	// Get Obtains the stored APIToken
	Get(ctx context.Context, value string) APIToken

	commonRepo
}

type IPRepository interface {
	commonRepo
}
