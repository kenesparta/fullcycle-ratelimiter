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

type APIKeyRedis struct {
	redisCli *redis.Client
}

func NewAPIKeyRedis(redisCli *redis.Client) *APIKeyRedis {
	return &APIKeyRedis{redisCli: redisCli}
}

func (at *APIKeyRedis) Save(ctx context.Context, token *entity.APIKey) (string, error) {
	req := dto.APIKeyInput{
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
		log.Println("error inserting API Key value")
		return "", redisErr
	}

	return token.Value(), nil
}

func (at *APIKeyRedis) Get(ctx context.Context, value string) entity.APIKey {
	return entity.APIKey{}
}

func (at *APIKeyRedis) UpsertRequest(ctx context.Context, key string, rl *entity.RateLimiter) error {
	return nil
}

func (at *APIKeyRedis) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	return nil
}

func (at *APIKeyRedis) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (at *APIKeyRedis) GetRequest(ctx context.Context, key string) (*entity.RateLimiter, error) {
	return nil, nil
}

func createAPKeyDurationPrefix(ip string) string {
	return fmt.Sprintf("%s:%s", entity.APIKeyPrefixDurationKey, ip)
}

func createAPIKeyRatePrefix(ip string) string {
	return fmt.Sprintf("%s:%s", entity.APIKeyPrefixRateKey, ip)
}
