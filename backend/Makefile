.PHONY: lillybox clean all test build

BUILD_DIR = build

BUILD_TARGET_BATCH = ./cmd/batch
BUILD_TARGET_HTTP = ./cmd/http
BUILD_TARGET_CHAT = ./cmd/chat

BUILD_OUTPUT_BATCH = batchd
BUILD_OUTPUT_HTTP = httpd
BUILD_OUTPUT_CHAT = chatd

all: build

build:
	@echo "\033[92mBuild batch daemon...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT_BATCH) $(BUILD_TARGET_BATCH)
	@echo "\033[92mBuild http daemon...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT_HTTP) $(BUILD_TARGET_HTTP)
	@echo "\033[92mBuild chat daemon...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT_CHAT) $(BUILD_TARGET_CHAT)

batch:
	@echo "\033[92mBuild batch daemon...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT_BATCH) $(BUILD_TARGET_BATCH)

http:
	@echo "\033[92mBuild http daemon...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT_HTTP) $(BUILD_TARGET_HTTP)

chat:
	@echo "\033[92mBuild chat daemon...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT_CHAT) $(BUILD_TARGET_CHAT)

clean:
	@echo "\033[92mClean daemon files...\033[0m"
	@rm $(BUILD_DIR)/*

fclean:
	@echo "\033[92mForce clean daemon files...\033[0m"
	@rm -f $(BUILD_DIR)/*

re:
	@make fclean
	@make all