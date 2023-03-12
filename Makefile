.DEFAULT_GOAL := test
SHELL := /bin/bash

# environment variables
BIN_DIR := $(CURDIR)/bin
GOBIN := $(BIN_DIR)
PATH := $(abspath $(BIN_DIR)):$(PATH)
export

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

.PHONY: test
test: ## run all tests
	go test ./... -race -cover

.PHONY: codegen_gql
codegen_gql: ## gqlgen code generate
	go run github.com/99designs/gqlgen generate

.PHONY: build
build:
	@CGO_ENABLED=0 go build -v -o $(BIN_DIR)/main ./cmd/main.go

.PHONY: run
run: build
	bin/main

.PHONY: up
up: ## start db
	@docker-compose up --build -d

.PHONY: down
down: ## stop db
	@docker-compose down

