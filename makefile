CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin

API_SRC=$(CURRENT_DIR)/cmd/bill-api
MAWINTER_SRC=$(CURRENT_DIR)/cmd/bill-mawinter

CONTAINER_NAME_API=bill-manager-api
CONTAINER_NAME_DB=bill-manager-db
CONTAINER_NAME_FETCHER=bill-manager-fetcher
CONTAINER_NAME_TWITTER=bill-manager-twitter

.PHONY: build start clean stop proto-build rebuild push
build:
	go build -a -tags "netgo" -installsuffix netgo  -ldflags="-s -w -extldflags \"-static\"" -o build/bin/ ./...
	docker build -t $(CONTAINER_NAME_API) -f build/dockerfile-api .
	docker build -t $(CONTAINER_NAME_DB) -f build/dockerfile-db .
	docker build -t $(CONTAINER_NAME_FETCHER) -f build/dockerfile-fetcher .
	docker build -t $(CONTAINER_NAME_TWITTER) -f build/dockerfile-twitter .

start:
	docker compose -f deploy/docker/docker-compose.yml up -d

stop:
	docker compose -f deploy/docker/docker-compose.yml down

clean:
	rm -rf build/bin/*

proto-build:
	protoc --go_out=. --go_opt=module=github.com/azuki774/bill-manager --go-grpc_out=. --go-grpc_opt=module=github.com/azuki774/bill-manager ./proto/*.proto
	python3 -m grpc_tools.protoc -I. --python_out=./fetcher/ --grpc_python_out=./fetcher/ ./proto/api.proto
	cp -rf fetcher/proto twclient/

rebuild:
	make stop && make clean && make && make start

push:	
	docker tag $(CONTAINER_NAME_API) ghcr.io/$(CONTAINER_NAME_API):develop
	docker tag $(CONTAINER_NAME_DB) ghcr.io/$(CONTAINER_NAME_DB):develop
	docker tag $(CONTAINER_NAME_FETCHER) ghcr.io/$(CONTAINER_NAME_FETCHER):develop
	docker tag $(CONTAINER_NAME_TWITTER) ghcr.io/$(CONTAINER_NAME_TWITTER):develop
	docker push ghcr.io/$(CONTAINER_NAME_API):develop
	docker push ghcr.io/$(CONTAINER_NAME_DB):develop
	docker push ghcr.io/$(CONTAINER_NAME_FETCHER):develop
	docker push ghcr.io/$(CONTAINER_NAME_TWITTER):develop
