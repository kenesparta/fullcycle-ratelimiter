SHELL := /bin/bash
-include .env

.PHONY: prepare
prepare:
	cp env.json.example env.json

.PHONY: init
init:
	go mod tidy

.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/api

.PHONY: run
run:
	docker compose -f docker-compose.yaml up -d --build

.PHONY: build-cli-test
build-cli-test:
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o cli-test ./cmd/cli