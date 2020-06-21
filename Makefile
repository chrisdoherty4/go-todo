
PKG_PATH := ./cmd/rememberme
BIN_NAME := rememberme

all:
	@go build -o $(BIN_NAME) $(PKG_PATH)

.PHONY: clean
clean: 
	@rm -f $(APP_NAME)

.PHONY: serve
serve:
	@$(PWD)/$(APP_NAME)
