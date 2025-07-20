# Go Echo App Makefile

# 变量定义
APP_NAME=go-echo-app
BUILD_DIR=build
MAIN_FILE=main.go

# 默认目标
.PHONY: help
help: ## 显示帮助信息
	@echo "可用的命令:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# 构建应用
.PHONY: build
build: ## 构建应用
	@echo "构建应用..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# 运行应用
.PHONY: run
run: ## 运行应用
	@echo "运行应用..."
	go run $(MAIN_FILE)

# 开发模式运行
.PHONY: dev
dev: ## 开发模式运行（自动重载）
	@echo "开发模式运行..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not found, installing..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

# 测试
.PHONY: test
test: ## 运行测试
	@echo "运行测试..."
	go test -v ./...

# 测试覆盖率
.PHONY: test-coverage
test-coverage: ## 运行测试并生成覆盖率报告
	@echo "运行测试覆盖率..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

# 清理
.PHONY: clean
clean: ## 清理构建文件
	@echo "清理构建文件..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# 格式化代码
.PHONY: fmt
fmt: ## 格式化代码
	@echo "格式化代码..."
	go fmt ./...

# 代码检查
.PHONY: lint
lint: ## 运行代码检查
	@echo "运行代码检查..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found, please install it first"; \
	fi

# 安装依赖
.PHONY: deps
deps: ## 安装依赖
	@echo "安装依赖..."
	go mod download
	go mod tidy

# 更新依赖
.PHONY: deps-update
deps-update: ## 更新依赖
	@echo "更新依赖..."
	go get -u ./...
	go mod tidy

# 生成文档
.PHONY: docs
docs: ## 生成API文档
	@echo "生成API文档..."
	@if command -v swag > /dev/null; then \
		swag init -g main.go; \
	else \
		echo "swag not found, please install it first: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

# Docker构建
.PHONY: docker-build
docker-build: ## 构建Docker镜像
	@echo "构建Docker镜像..."
	docker build -t $(APP_NAME) .

# Docker运行
.PHONY: docker-run
docker-run: ## 运行Docker容器
	@echo "运行Docker容器..."
	docker run -p 8080:8080 $(APP_NAME)

# 安装开发工具
.PHONY: install-tools
install-tools: ## 安装开发工具
	@echo "安装开发工具..."
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/swaggo/swag/cmd/swag@latest

# 创建新版本
.PHONY: release
release: ## 创建发布版本
	@echo "创建发布版本..."
	@read -p "Enter version (e.g., v1.0.0): " version; \
	git tag $$version; \
	git push origin $$version; \
	echo "Release $$version created and pushed"

# 数据库迁移
.PHONY: migrate
migrate: ## 运行数据库迁移
	@echo "运行数据库迁移..."
	@echo "请实现数据库迁移逻辑"

# 种子数据
.PHONY: seed
seed: ## 插入种子数据
	@echo "插入种子数据..."
	@echo "请实现种子数据插入逻辑"