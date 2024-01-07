package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kenesparta/fullcycle-ratelimiter/config"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/dto"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/database"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/usecase"
	"github.com/redis/go-redis/v9"
)

type Middleware struct {
	RedisClient *redis.Client
	Config      *config.Config
}

func (a *Middleware) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ipDB := database.NewIPRedis(a.RedisClient)
			ipReq := usecase.NewRegisterIPUseCase(ipDB, a.Config)
			execute, execErr := ipReq.Execute(r.Context(), dto.IPRequestSave{
				IP: func() string {
					spltStr := strings.Split(r.RemoteAddr, ":")
					if len(spltStr) > 0 {
						return spltStr[0]
					}
					return ""
				}(),
				TimeAdd: time.Now(),
			})
			if execErr != nil {
				log.Printf("Error executing NewRegisterIPUseCase: %s\n", execErr.Error())
				http.Error(w, execErr.Error(), http.StatusInternalServerError)
				return
			}

			if !execute.Allow {
				log.Printf("Too many request: %s\n", entity.ErrExceededRequest.Error())
				http.Error(w, entity.ErrExceededRequest.Error(), http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
