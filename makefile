CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin

API_SRC=$(CURRENT_DIR)/cmd/bill-api
MAWINTER_SRC=$(CURRENT_DIR)/cmd/bill-mawinter
TWITTER_SRC=$(CURRENT_DIR)/cmd/bill-twitter

.PHONY: build run clean stop
build:
	cd $(API_SRC) && CGO_ENABLED=0 go build -o $(BIN_DIR)/bill-manager-api
	cd $(MAWINTER_SRC) && CGO_ENABLED=0 go build -o $(BIN_DIR)/bill-manager-mawinter
	cd $(TWITTER_SRC) && CGO_ENABLED=0 go build -o $(BIN_DIR)/bill-manager-twitter

	docker build -t azuki774/bill-manager-fetcher -f build/dockerfile-fetcher .
	docker build -t azuki774/bill-manager-api -f build/dockerfile-api .
	docker build -t azuki774/bill-manager-db -f build/dockerfile-db .
	docker build -t azuki774/bill-manager-mawinter -f build/dockerfile-mawinter .
	docker build -t azuki774/bill-manager-twitter -f build/dockerfile-twitter .

clean:
	rm -rf build/bin/*
