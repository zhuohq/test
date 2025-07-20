# Go Echo Web Framework Project

一个基于 Echo 框架的标准 Go Web 应用程序模板，包含完整的项目结构和最佳实践。

## 🚀 特性

- **Echo Web 框架**: 高性能、可扩展的 Go Web 框架
- **标准项目结构**: 遵循 Go 项目最佳实践
- **中间件支持**: 内置日志、CORS、恢复等中间件
- **配置管理**: 环境变量配置系统
- **统一响应格式**: 标准化的 API 响应结构
- **数据验证**: 请求数据验证和错误处理
- **Docker 支持**: 完整的容器化部署方案
- **开发工具**: Makefile 提供常用命令

## 📁 项目结构

```
.
├── main.go                 # 应用程序入口
├── go.mod                  # Go 模块文件
├── go.sum                  # 依赖校验文件
├── .env.example           # 环境变量示例
├── Dockerfile             # Docker 构建文件
├── docker-compose.yml     # Docker Compose 配置
├── Makefile               # 构建和部署脚本
├── README.md              # 项目文档
├── internal/              # 内部包
│   ├── config/           # 配置管理
│   ├── handlers/         # HTTP 处理器
│   ├── middleware/       # 自定义中间件
│   ├── models/           # 数据模型
│   └── routes/           # 路由配置
└── pkg/                  # 公共包
    ├── utils/            # 工具函数
    └── validator/        # 数据验证
```

## 🛠️ 安装和运行

### 前置要求

- Go 1.21 或更高版本
- Docker 和 Docker Compose (可选)

### 快速开始

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd go-echo-app
   ```

2. **安装依赖**
   ```bash
   make deps
   # 或者
   go mod download
   go mod tidy
   ```

3. **配置环境变量**
   ```bash
   cp .env.example .env
   # 编辑 .env 文件
   ```

4. **运行应用**
   ```bash
   # 开发模式
   make dev
   
   # 或者直接运行
   make run
   ```

5. **访问应用**
   - 主页: http://localhost:8080
   - 健康检查: http://localhost:8080/health
   - API 文档: http://localhost:8080/api/v1/users

### 使用 Docker

1. **构建和运行**
   ```bash
   # 使用 Docker Compose
   docker-compose up -d
   
   # 或者单独构建
   make docker-build
   make docker-run
   ```

2. **停止服务**
   ```bash
   docker-compose down
   ```

## 📚 API 文档

### 基础端点

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/` | 应用首页 |
| GET | `/health` | 健康检查 |

### 用户 API

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/v1/users` | 获取所有用户 |
| GET | `/api/v1/users/:id` | 获取指定用户 |
| POST | `/api/v1/users` | 创建新用户 |
| PUT | `/api/v1/users/:id` | 更新用户 |
| DELETE | `/api/v1/users/:id` | 删除用户 |

### 请求示例

**创建用户**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

**获取用户列表**
```bash
curl http://localhost:8080/api/v1/users
```

## 🛠️ 开发

### 可用命令

```bash
# 显示所有可用命令
make help

# 开发模式运行（自动重载）
make dev

# 运行测试
make test

# 代码格式化
make fmt

# 代码检查
make lint

# 构建应用
make build

# 清理构建文件
make clean

# 安装开发工具
make install-tools
```

### 添加新功能

1. **添加新的处理器**
   - 在 `internal/handlers/` 目录下创建新文件
   - 实现处理函数

2. **添加新的路由**
   - 在 `internal/routes/routes.go` 中添加路由配置

3. **添加新的中间件**
   - 在 `internal/middleware/` 目录下创建新文件

4. **添加新的模型**
   - 在 `internal/models/` 目录下创建新文件

## 🔧 配置

### 环境变量

| 变量名 | 默认值 | 描述 |
|--------|--------|------|
| `PORT` | `8080` | 服务器端口 |
| `HOST` | `localhost` | 服务器主机 |
| `DB_HOST` | `localhost` | 数据库主机 |
| `DB_PORT` | `5432` | 数据库端口 |
| `DB_USER` | `postgres` | 数据库用户 |
| `DB_PASSWORD` | `` | 数据库密码 |
| `DB_NAME` | `go_echo_app` | 数据库名称 |
| `JWT_SECRET` | `your-secret-key` | JWT 密钥 |

## 🧪 测试

```bash
# 运行所有测试
make test

# 运行测试并生成覆盖率报告
make test-coverage
```

## 📦 部署

### 生产环境部署

1. **使用 Docker**
   ```bash
   docker-compose -f docker-compose.prod.yml up -d
   ```

2. **直接部署**
   ```bash
   make build
   ./build/go-echo-app
   ```

### 环境配置

- 复制 `.env.example` 到 `.env`
- 修改生产环境配置
- 设置适当的数据库连接
- 配置 JWT 密钥

## 🤝 贡献

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- [Echo Framework](https://echo.labstack.com/)
- [Go Programming Language](https://golang.org/)
- [Docker](https://www.docker.com/)

## 📞 支持

如果您遇到任何问题或有疑问，请：

1. 查看 [Issues](../../issues)
2. 创建新的 Issue
3. 联系维护者

---

**Happy Coding! 🎉**