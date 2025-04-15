# 定义变量
BIN_DIR := ./bin
APP_NAME := safelineApi
SRC_DIR := ./cmd/safelineApi
VERSION := 1.0.0
BUILD_TIME := $(shell date +"%Y-%m-%dT%H:%M:%S")

# 默认任务
.DEFAULT_GOAL := build

# 构建任务
build:
	@echo "Building $(APP_NAME) version $(VERSION)..."
	mkdir -p $(BIN_DIR)
	go build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o $(BIN_DIR)/$(APP_NAME) $(SRC_DIR)

# 运行任务
run:
	@echo "Running $(APP_NAME)..."
	$(BIN_DIR)/$(APP_NAME)

# 清理任务
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)

# 测试任务
test:
	@echo "Running tests..."
	go test ./...

# 格式化代码
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 检查代码风格
vet:
	@echo "Vetting code..."
	go vet ./...

# 安装依赖
tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# 多平台编译
build-all:
	@echo "Building for all platforms..."
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)-linux-amd64 $(SRC_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)-windows-amd64.exe $(SRC_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)-darwin-amd64 $(SRC_DIR)