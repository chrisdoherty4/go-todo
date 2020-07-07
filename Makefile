
PKG_PATH := ./cmd/rememberme
BIN_NAME := rememberme
BIN_PATH := $(PWD)/$(BIN_NAME)

all: test build

.PHONY: build
build:
	@go build -o $(BIN_PATH) $(PKG_PATH)

.PHONY: clean
clean: 
	@rm -f $(BIN_PATH)

.PHONY: serve
serve:
	@$(BIN_PATH)

.PHONY: test
test:
	@go test -v ./...