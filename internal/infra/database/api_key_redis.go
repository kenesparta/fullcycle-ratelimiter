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
	"time"
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
	req := dto.APIKeyRequestDB{
		MaxRequests:   rl.MaxRequests,
		TimeWindowSec: rl.TimeWindowSec,
		Requests: func() []int64 {
			reqInt := make([]int64, 0)
			for _, r := range rl.Requests {
				reqInt = append(reqInt, r.Unix())
			}
			return reqInt
		}(),
	}

	jsonReq, marErr := json.Marshal(req)
	if marErr != nil {
		log.Println("error marshaling API Key")
		return marErr
	}

	redisErr := at.redisCli.Set(ctx, createAPIKeyRatePrefix(key), jsonReq, 0).Err()
	if redisErr != nil {
		log.Println("error inserting API Key value")
		return redisErr
	}

	return nil
}

func (at *APIKeyRedis) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	if redisErr := at.redisCli.Set(
		ctx,
		createAPIKeyDurationPrefix(key),
		entity.StatusIPBlocked,
		time.Second*time.Duration(BlockedDuration),
	).Err(); redisErr != nil {
		log.Println("error inserting SaveBlockedDuration on API Key")
		return redisErr
	}

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

// GetRequest reads the stored array of request
func (at *APIKeyRedis) GetRequest(ctx context.Context, key string) (*entity.RateLimiter, error) {
	val, getErr := at.redisCli.Get(ctx, createAPIKeyRatePrefix(key)).Result()
	if errors.Is(getErr, redis.Nil) {
		log.Println("INFO: GetRequest API key does not exist")
		return &entity.RateLimiter{
			Requests:      make([]time.Time, 0),
			TimeWindowSec: 0,
			MaxRequests:   0,
		}, nil
	}
	if getErr != nil {
		return nil, getErr
	}

	var rateLimiter dto.APIKeyRequestDB
	if err := json.Unmarshal([]byte(val), &rateLimiter); err != nil {
		log.Println("API key RateLimiter unmarshal error")
		return &entity.RateLimiter{}, err
	}

	return &entity.RateLimiter{
		Requests: func() []time.Time {
			reqTimeStamp := make([]time.Time, 0)
			for _, rr := range rateLimiter.Requests {
				reqTimeStamp = append(reqTimeStamp, time.Unix(rr, 0))
			}
			return reqTimeStamp
		}(),
		TimeWindowSec: rateLimiter.TimeWindowSec,
		MaxRequests:   rateLimiter.MaxRequests,
	}, nil
}

func createAPIKeyDurationPrefix(key string) string {
	return fmt.Sprintf("%s:%s", entity.APIKeyPrefixDurationKey, key)
}

func createAPIKeyRatePrefix(key string) string {
	return fmt.Sprintf("%s:%s", entity.APIKeyPrefixRateKey, key)
}
