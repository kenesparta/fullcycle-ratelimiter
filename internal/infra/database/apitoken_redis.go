package database

import (
	"context"

	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	"github.com/redis/go-redis/v9"
)

type IPRedis struct {
	redisCli *redis.Client
}

func NewIPRedis(redisCli *redis.Client) *IPRedis {
	return &IPRedis{redisCli: redisCli}
}

func (ip *IPRedis) UpsertRequest(ctx context.Context, key string, rl *entity.RateLimiter) error {
	return nil
}

// SaveBlockedDuration Stores the blocked duration amount by key
func (ip *IPRedis) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	return nil
}

// GetBlockedDuration Obtain the blocked duration by key
func (ip *IPRedis) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	return "", nil
}

// GetRequest reads the stored array of request
func (ip *IPRedis) GetRequest(ctx context.Context, value string) (*entity.RateLimiter, error) {
	return nil, nil
}
