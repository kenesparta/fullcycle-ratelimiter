package middleware

import (
	"errors"
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

type IPMiddleware struct {
	RedisClient *redis.Client
	Config      *config.Config
}

func getIP(remoteAddr string) string {
	splitStr := strings.Split(remoteAddr, ":")
	if len(splitStr) > 0 {
		return splitStr[0]
	}

	return ""
}

func (ip *IPMiddleware) Execute(w http.ResponseWriter, r *http.Request) error {
	ipDB := database.NewIPRedis(ip.RedisClient)
	ipReq := usecase.NewRegisterIPUseCase(ipDB, ip.Config)
	execute, execErr := ipReq.Execute(r.Context(), dto.IPRequestSave{
		IP:      getIP(r.RemoteAddr),
		TimeAdd: time.Now(),
	})
	if errors.Is(execErr, entity.ErrIPExceededAmountRequest) {
		log.Printf("Error executing NewRegisterIPUseCase: %s\n", execErr.Error())
		http.Error(w, execErr.Error(), http.StatusTooManyRequests)
		return execErr
	}
	if execErr != nil {
		log.Printf("Error executing NewRegisterIPUseCase: %s\n", execErr.Error())
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
		return execErr
	}

	if !execute.Allow {
		log.Printf("Too many request: %s\n", entity.ErrExceededRequest.Error())
		http.Error(w, entity.ErrExceededRequest.Error(), http.StatusTooManyRequests)
		return errors.New("too many request")
	}

	return nil
}
