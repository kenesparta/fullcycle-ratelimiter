package main

import (
	"log"
	"net/http"
	
	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/webserver"
)

func main() {
	cfg, _ := NewRunConfig()
	ipHandler := NewIPHandler()
	newWebServer := webserver.NewWebServer(cfg.Config.App.Port)
	newWebServer.AddHandler(http.MethodGet, "/hello-world", ipHandler.HelloWorld)
	log.Println("Starting web server on port", cfg.Config.App.Port)
	newWebServer.Start()
}
