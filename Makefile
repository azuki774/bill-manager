COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''
SHELL=/bin/bash

CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin

API_SRC=$(CURRENT_DIR)/cmd/bill-api
MAWINTER_SRC=$(CURRENT_DIR)/cmd/bill-mawinter

CONTAINER_NAME_REMIX=bill-manager-remix
CONTAINER_NAME_TWITTER=bill-manager-twitter
CONTAINER_NAME_MAWINTER=bill-manager-mawinter

.PHONY: build start clean stop test rebuild
build:
	go build -a -tags "netgo" -installsuffix netgo  -ldflags="-s -w -extldflags \"-static\"" -o build/bin/ ./...
	docker build -t $(CONTAINER_NAME_REMIX) -f build/dockerfile-remix .
	docker build -t $(CONTAINER_NAME_TWITTER) -f build/dockerfile-twitter .
	docker build -t $(CONTAINER_NAME_MAWINTER) -f build/dockerfile-mawinter .

start:
	docker compose -f deployment/compose.yml up -d

stop:
	docker compose -f deployment/compose.yml down

clean:
	rm -rf build/bin/*

test:
	gofmt -l .
	go vet ./...
	staticcheck ./...
	go test -v ./...  | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

rebuild:
	make stop && make clean && make && make start