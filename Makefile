# Project: Go E-Shop Service
APP_NAME := go-e-shop-service
VERSION := 1.0.0
BUILD_DIR := build
SRC_DIR := ./cmd/server/main.go
SRC_MIGRATION_DIR := ./cmd/migration
BINARY := $(BUILD_DIR)/$(APP_NAME)
GO_FILES := $(shell find . -type f -name '*.go')

.PHONY: all build clean run test

all: build

build: $(BINARY)

$(BINARY): $(GO_FILES)
	mkdir -p $(BUILD_DIR)
	go build -o $(BINARY) $(SRC_DIR)

clean:
	rm -rf $(BUILD_DIR)

run: build
	$(BINARY)

test:
	go test ./... -v

migrate-up: 
	go run $(SRC_MIGRATION_DIR) up

migrate-down: 
	go run $(SRC_MIGRATION_DIR) down
