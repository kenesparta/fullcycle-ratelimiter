package main

import (
	"github.com/kenesparta/fullcycle-ratelimiter/internal/infra/handler"
)

func NewIPHandler() *handler.IPHandler {
	return handler.NewIPHandler(nil)
}
