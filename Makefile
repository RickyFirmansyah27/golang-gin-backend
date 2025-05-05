# Variables
APP_NAME=app
BUILD_DIR=build

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

# Main targets
.PHONY: all build run test clean

all: clean build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) ./$(APP_NAME)/main.go

start:
	$(BUILD_DIR)/$(APP_NAME)

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Development targets
.PHONY: dev hot-reload

run:
	$(GORUN) ./$(APP_NAME)/main.go

live:
	air

# Dependencies
.PHONY: deps

deps:
	$(GOCMD) mod tidy
	$(GOCMD) mod download

# Docker targets
.PHONY: docker-build docker-run

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run -p 8080:8080 $(APP_NAME)
