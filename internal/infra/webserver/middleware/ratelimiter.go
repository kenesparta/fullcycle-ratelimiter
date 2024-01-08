package middleware

import (
	"github.com/kenesparta/fullcycle-ratelimiter/config"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type Middleware struct {
	RedisClient *redis.Client
	Config      *config.Config
}

func (m *Middleware) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get(entity.APIKeyHeaderName)
			strategy := Factory(apiKey, m)
			if err := strategy.Execute(w, r); err != nil {
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
