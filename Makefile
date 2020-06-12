
APP_NAME = winston

all:
	@go build -o $(APP_NAME) .

.PHONY: clean
clean: 
	@rm $(APP_NAME)

.PHONY: serve
serve:
	@$(PWD)/$(APP_NAME)
