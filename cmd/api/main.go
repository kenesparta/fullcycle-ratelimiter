package main

import (
	"fmt"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/webserver/middleware"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"

	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/webserver"
)

func main() {
	cfg, _ := NewRunConfig()
	redisCli := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%s", cfg.Config.Redis.Host, cfg.Config.Redis.Port),
			DB:   cfg.Config.Redis.Db,
		},
	)

	ipHandler := NewIPHandler()
	newWebServer := webserver.NewWebServer(cfg.Config.App.Port)
	newWebServer.InternalMiddleware = middleware.Middleware{
		RedisClient: redisCli,
		Config:      cfg.Config,
	}

	newWebServer.AddHandler(http.MethodGet, "/hello-world", ipHandler.HelloWorld)
	log.Println("Starting web server on port", cfg.Config.App.Port)
	newWebServer.Start()
}
