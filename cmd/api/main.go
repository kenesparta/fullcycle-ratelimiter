package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/database"
	internalHandler "github.com/kenesparta/fullcycle-ratelimiter/internal/infra/handler"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/webserver"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/webserver/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	cfg, _ := NewRunConfig()
	redisCli := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%s", cfg.Config.Redis.Host, cfg.Config.Redis.Port),
			DB:   cfg.Config.Redis.Db,
		},
	)

	newWebServer := webserver.NewWebServer(cfg.Config.App.Port)
	newWebServer.InternalMiddleware = middleware.Middleware{
		RedisClient: redisCli,
		Config:      cfg.Config,
	}
	apikeyHandler := internalHandler.NewAPIKeyHandler(database.NewAPIKeyRedis(redisCli))

	newWebServer.AddHandler(http.MethodGet, "/hello-world", internalHandler.HelloWorld)
	newWebServer.AddHandler(http.MethodPost, "/api-key", apikeyHandler.CreateToken)
	log.Println("Starting web server on port", cfg.Config.App.Port)
	newWebServer.Start()
}
