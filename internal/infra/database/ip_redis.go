package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
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
	req := dto.IPRequestDB{
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
		log.Println("error marshaling IP")
		return marErr
	}

	redisErr := ip.redisCli.Set(ctx, createIPRatePrefix(key), jsonReq, 0).Err()
	if redisErr != nil {
		log.Println("error inserting IP value")
		return redisErr
	}

	return nil
}

// SaveBlockedDuration Stores the blocked duration amount by key
func (ip *IPRedis) SaveBlockedDuration(ctx context.Context, key string, BlockedDuration int64) error {
	if redisErr := ip.redisCli.Set(
		ctx,
		createIPDurationPrefix(key),
		entity.StatusIPBlocked,
		time.Second*time.Duration(BlockedDuration),
	).Err(); redisErr != nil {
		log.Println("error inserting SaveBlockedDuration for IP")
		return redisErr
	}

	return nil
}

// GetBlockedDuration Obtain the blocked duration by key
func (ip *IPRedis) GetBlockedDuration(ctx context.Context, key string) (string, error) {
	val, getErr := ip.redisCli.Get(ctx, createIPDurationPrefix(key)).Result()
	if errors.Is(getErr, redis.Nil) {
		log.Println("INFO: GetBlockedDuration IP key does not exist")
		return "", nil
	}
	if getErr != nil {
		return "", getErr
	}

	return val, nil
}

// GetRequest reads the stored array of request
func (ip *IPRedis) GetRequest(ctx context.Context, key string) (*entity.RateLimiter, error) {
	val, getErr := ip.redisCli.Get(ctx, createIPRatePrefix(key)).Result()
	if errors.Is(getErr, redis.Nil) {
		log.Println("INFO: GetRequest IP key does not exist")
		return &entity.RateLimiter{
			Requests:      make([]time.Time, 0),
			TimeWindowSec: 0,
			MaxRequests:   0,
		}, nil
	}
	if getErr != nil {
		return nil, getErr
	}

	var rateLimiter dto.IPRequestDB
	if err := json.Unmarshal([]byte(val), &rateLimiter); err != nil {
		log.Println("IP RateLimiter unmarshal error")
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

func createIPDurationPrefix(ip string) string {
	return fmt.Sprintf("%s:%s", entity.IPPrefixDurationKey, ip)
}

func createIPRatePrefix(ip string) string {
	return fmt.Sprintf("%s:%s", entity.IPPrefixRateKey, ip)
}
