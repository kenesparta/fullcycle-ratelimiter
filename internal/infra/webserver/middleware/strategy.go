package middleware

import (
	"net/http"
)

type StrategyMiddleware interface {
	Execute(w http.ResponseWriter, r *http.Request) error
}

func Factory(apiKey string, m *Middleware) StrategyMiddleware {
	if apiKey != "" {
		return &APIKeyMiddleware{RedisClient: m.RedisClient, ApiKey: apiKey}
	}

	return &IPMiddleware{RedisClient: m.RedisClient, Config: m.Config}
}
