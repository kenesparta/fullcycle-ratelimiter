package middleware

import (
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	"net/http"
)

type StrategyMiddleware interface {
	Execute(w http.ResponseWriter, r *http.Request) error
}

func Factory(apiKey string, m *Middleware) StrategyMiddleware {
	switch apiKey {
	case entity.APIKeyHeaderName:
		return &APIKeyMiddleware{RedisClient: m.RedisClient, ApiKey: apiKey}
	default:
		return &IPMiddleware{RedisClient: m.RedisClient, Config: m.Config}
	}
}
