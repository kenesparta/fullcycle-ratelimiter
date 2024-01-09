package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	internalMw "github.com/kenesparta/fullcycle-ratelimiter/internal/infra/webserver/middleware"
)

type HandlerProps struct {
	Method string
	Path   string
	Func   http.HandlerFunc
}

type WebServer struct {
	WebServerPort      string
	Router             chi.Router
	Handlers           []HandlerProps
	InternalMiddleware internalMw.Middleware
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]HandlerProps, 0),
		WebServerPort: fmt.Sprintf("0.0.0.0:%s", serverPort),
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, HandlerProps{
		Method: method,
		Path:   path,
		Func:   handler,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(s.InternalMiddleware.RateLimiter)
	for _, h := range s.Handlers {
		s.Router.Method(h.Method, h.Path, h.Func)
	}

	if err := http.ListenAndServe(s.WebServerPort, s.Router); err != nil {
		log.Printf("error starting the server.")
		return
	}
}
