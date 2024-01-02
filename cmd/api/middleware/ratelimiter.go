package middleware

import (
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
)

const RedisContextName = "redis-cli"

type Middleware struct {
	RedisClient *redis.Client
}

func (a *Middleware) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RedisContextName, a.RedisClient)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
