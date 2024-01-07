package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	"github.com/redis/go-redis/v9"
	"log"
)

type APITokenRedis struct {
	redisCli *redis.Client
}

func NewAPITokenRedis(redisCli *redis.Client) *APITokenRedis {
	return &APITokenRedis{redisCli: redisCli}
}

func (at *APITokenRedis) Save(ctx context.Context, token *entity.APIToken) (string, error) {
	req := dto.APITokenInput{
		MaxRequests:     token.RateLimiter.MaxRequests,
		TimeWindowSec:   token.RateLimiter.TimeWindowSec,
		BlockedDuration: token.BlockedDuration,
	}

	jsonReq, marErr := json.Marshal(req)
	if marErr != nil {
		log.Println("error marshaling API Token")
		return "", marErr
	}

	if redisErr := at.redisCli.Set(
		ctx,
		token.Value(),
		jsonReq,
		0,
	).Err(); redisErr != nil {
		log.Println("error inserting API Token value")
		return "", redisErr
	}

	return token.Value(), nil
}

func (at *APITokenRedis) Get(ctx context.Context, value string) entity.APIToken {
	return entity.APIToken{}
}

func (at *APITokenRedis) UpsertRequest(ctx context.Context, key string, rl *entity.RateLimiter) error {
	return nil
}

func (at *APITokenRedis) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	return nil
}

func (at *APITokenRedis) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (at *APITokenRedis) GetRequest(ctx context.Context, key string) (*entity.RateLimiter, error) {
	return nil, nil
}

func createAPITokenDurationPrefix(ip string) string {
	return fmt.Sprintf("%s:%s", entity.APITokenPrefixDurationKey, ip)
}

func createAPITokenRatePrefix(ip string) string {
	return fmt.Sprintf("%s:%s", entity.APITokenPrefixRateKey, ip)
}
