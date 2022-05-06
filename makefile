CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin

API_SRC=$(CURRENT_DIR)/cmd/bill-api
MAWINTER_SRC=$(CURRENT_DIR)/cmd/bill-mawinter

.PHONY: build run clean stop proto-build
build:
	cd $(API_SRC) && CGO_ENABLED=0 go build -o $(BIN_DIR)/bill-manager-api
	cd $(MAWINTER_SRC) && CGO_ENABLED=0 go build -o $(BIN_DIR)/bill-manager-mawinter

	docker build -t azuki774/bill-manager-fetcher -f build/dockerfile-fetcher .
	docker build -t azuki774/bill-manager-api -f build/dockerfile-api .
	docker build -t azuki774/bill-manager-db -f build/dockerfile-db .
	docker build -t azuki774/bill-manager-mawinter -f build/dockerfile-mawinter .
	docker build -t azuki774/bill-manager-twitter -f build/dockerfile-twitter .

run:
	docker-compose -f deploy/docker/docker-compose.yml up -d

stop:
	docker-compose -f deploy/docker/docker-compose.yml down

clean:
	rm -rf build/bin/*

proto-build:
	protoc --go_out=. --go_opt=module=github.com/azuki774/bill-manager --go-grpc_out=. --go-grpc_opt=module=github.com/azuki774/bill-manager ./proto/*.proto
	python3 -m grpc_tools.protoc -I. --python_out=./fetcher/ --grpc_python_out=./fetcher/ ./proto/api.proto
	cp -rf fetcher/proto twclient/proto
