package main

import (
	"github.com/kenesparta/fullcycle-ratelimiter/config"
	"github.com/kenesparta/fullcycle-ratelimiter/config/impl"
)

const ConfigFileName = "env"

type RunConfig struct {
	Config *config.Config
}

func NewRunConfig() (*RunConfig, error) {
	cfg, err := config.NewConfig(impl.NewViper(ConfigFileName))
	if err != nil {
		return nil, err
	}

	return &RunConfig{
		Config: cfg,
	}, nil
}
