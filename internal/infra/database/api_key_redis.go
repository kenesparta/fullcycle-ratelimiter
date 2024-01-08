package database

import (
	"context"
	"encoding/json"
	"errors"
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

func (at *APIKeyRedis) Save(ctx context.Context, key *entity.APIKey) (string, error) {
	req := dto.APIKeyInput{
		MaxRequests:     key.RateLimiter.MaxRequests,
		TimeWindowSec:   key.RateLimiter.TimeWindowSec,
		BlockedDuration: key.BlockedDuration,
	}

	jsonReq, marErr := json.Marshal(req)
	if marErr != nil {
		log.Println("error marshaling API Key")
		return "", marErr
	}

	if redisErr := at.redisCli.Set(
		ctx,
		key.Value(),
		jsonReq,
		0,
	).Err(); redisErr != nil {
		log.Println("error inserting API Key value")
		return "", redisErr
	}

	return key.Value(), nil
}

func (at *APIKeyRedis) Get(ctx context.Context, value string) (*entity.APIKey, error) {
	val, getErr := at.redisCli.Get(ctx, value).Result()
	if errors.Is(getErr, redis.Nil) {
		log.Println("API key does not exist")
		return &entity.APIKey{}, getErr
	}
	if getErr != nil {
		return &entity.APIKey{}, getErr
	}

	var apiKeyConfigDB dto.APIKeyInput
	if err := json.Unmarshal([]byte(val), &apiKeyConfigDB); err != nil {
		log.Println("API key configuration marshall error")
		return &entity.APIKey{}, err
	}

	return &entity.APIKey{
		BlockedDuration: apiKeyConfigDB.BlockedDuration,
		RateLimiter: entity.RateLimiter{
			TimeWindowSec: apiKeyConfigDB.TimeWindowSec,
			MaxRequests:   apiKeyConfigDB.MaxRequests,
		},
	}, nil
}

func (at *APIKeyRedis) UpsertRequest(ctx context.Context, key string, rl *entity.RateLimiter) error {
	return nil
}

func (at *APIKeyRedis) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	return nil
}

func (at *APIKeyRedis) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	val, getErr := at.redisCli.Get(ctx, createAPIKeyDurationPrefix(key)).Result()
	if errors.Is(getErr, redis.Nil) {
		log.Println("API key does not exist")
		return "", nil
	}
	if getErr != nil {
		return "", getErr
	}

	return val, nil
}

func (at *APIKeyRedis) GetRequest(ctx context.Context, key string) (*entity.RateLimiter, error) {
	return nil, nil
}

func createAPIKeyDurationPrefix(ip string) string {
	return fmt.Sprintf("%s:%s", entity.APIKeyPrefixDurationKey, ip)
}

func createAPIKeyRatePrefix(ip string) string {
	return fmt.Sprintf("%s:%s", entity.APIKeyPrefixRateKey, ip)
}
