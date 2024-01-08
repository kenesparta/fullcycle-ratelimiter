package main

import (
	"net/http"

	"github.com/kenesparta/fullcycle-ratelimiter/config"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/database"
	internalHandler "github.com/kenesparta/fullcycle-ratelimiter/internal/infra/handler"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/webserver"
	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/webserver/middleware"
	"github.com/redis/go-redis/v9"
)

func CreateWebServer(cfg *config.Config, redisCli *redis.Client) *webserver.WebServer {
	newWebServer := webserver.NewWebServer(cfg.App.Port)
	newWebServer.InternalMiddleware = middleware.Middleware{
		RedisClient: redisCli,
		Config:      cfg,
	}
	apikeyHandler := internalHandler.NewAPIKeyHandler(database.NewAPIKeyRedis(redisCli))

	newWebServer.AddHandler(http.MethodPost, "/api-key", apikeyHandler.CreateAPIKey)
	newWebServer.AddHandler(http.MethodGet, "/hello-world", internalHandler.HelloWorld)
	newWebServer.AddHandler(http.MethodGet, "/hello-world-key", internalHandler.HelloWorldWithAPIKey)

	return newWebServer
}
