#! /usr/bin/make -f

# ---------------------------------------------------------
# Project variables.
# ---------------------------------------------------------
PROJECT_NAME := $(shell basename "$(PWD)")
PACKAGE := github.com/knstl/$(PROJECT_NAME)
VERSION := $(shell git describe --tags 2>/dev/null || git describe --all)
BUILD := $(shell git rev-parse --short HEAD)
DATETIME := $(shell date +"%Y.%m.%d-%H:%M:%S")

# ---------------------------------------------------------
# Use linker flags to provide version/build settings.
# ---------------------------------------------------------
LDFLAGS=-ldflags "-X=$(PACKAGE)/build.Version=$(VERSION) -X=$(PACKAGE)/build.Build=$(BUILD) -X=$(PACKAGE)/build.Date=$(DATETIME)"

# ---------------------------------------------------------
# Go related variables.
# ---------------------------------------------------------
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

# ---------------------------------------------------------
# Go files.
# ---------------------------------------------------------
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

# ---------------------------------------------------------
# Docker Variables
# ---------------------------------------------------------
DOCKER_COMPOSE_FILE_LOCAL_ENV ?= ./docker-compose.yaml

# ---------------------------------------------------------
# Code Checking
# ---------------------------------------------------------
check: fmt lint unit-test

unit-test:
	@echo "  >  Running unit tests"
	go clean -testcache
	GOBIN=$(GOBIN) go test -cover -race -coverprofile=coverage.txt -covermode=atomic -v ./...

fmt:
	@echo "  >  Format all go files"
	GOBIN=$(GOBIN) gofmt -w ${GOFMT_FILES}

lint-install:
ifeq ("$(wildcard bin/golangci-lint)","")
	@echo "  >  Installing golint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s
endif

lint: lint-install
	@echo "  >  Running golint"
	bin/golangci-lint run --timeout=2m