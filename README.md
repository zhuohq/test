# Go Echo API

一个基于Echo框架的Go Web API项目。

## 功能特性

- RESTful API设计
- 用户管理CRUD操作
- 健康检查端点
- CORS支持
- 日志记录
- 错误恢复

## 系统要求

- Go 1.21 或更高版本
- Windows/Linux/macOS

## 安装和运行

### 方法1: 直接运行（推荐）

#### Windows
```bash
# 使用批处理文件
run.bat

# 或使用PowerShell
.\run.ps1

# 或直接运行
go mod download
go mod tidy
go run main.go
```

#### Linux/macOS
```bash
# 使用Makefile
make run

# 或直接运行
go mod download
go mod tidy
go run main.go
```

### 方法2: 构建后运行

```bash
# 构建应用
go build -o app main.go

# 运行构建后的应用
./app  # Linux/macOS
app.exe  # Windows
```

## API端点

### 健康检查
- `GET /health` - 服务器状态检查

### 用户管理
- `GET /api/v1/users` - 获取所有用户
- `GET /api/v1/users/:id` - 获取指定用户
- `POST /api/v1/users` - 创建新用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 根路径
- `GET /` - API信息

## 开发

### 安装依赖
```bash
go mod download
go mod tidy
```

### 运行测试
```bash
go test -v ./...
```

### 代码格式化
```bash
go fmt ./...
```

## Docker支持

### 构建镜像
```bash
docker build -t go-echo-app .
```

### 运行容器
```bash
docker run -p 8080:8080 go-echo-app
```

## 故障排除

### Windows下常见问题

1. **依赖下载失败**
   ```bash
   go mod download
   go mod tidy
   ```

2. **端口被占用**
   - 检查8080端口是否被其他程序占用
   - 可以修改main.go中的端口号

3. **Go未安装或不在PATH中**
   - 下载并安装Go: https://golang.org/dl/
   - 确保Go在系统PATH中

4. **权限问题**
   - 以管理员身份运行命令提示符或PowerShell

### 常见错误解决

- `missing go.sum entry`: 运行 `go mod tidy`
- `module not found`: 运行 `go mod download`
- `port already in use`: 更改端口号或停止占用端口的程序

## 项目结构

```
.
├── main.go              # 主程序入口
├── go.mod               # Go模块文件
├── go.sum               # 依赖校验文件
├── run.bat              # Windows批处理启动脚本
├── run.ps1              # PowerShell启动脚本
├── Makefile             # Linux/macOS构建脚本
├── Dockerfile           # Docker配置
├── README.md            # 项目说明
└── internal/            # 内部包
    ├── config/          # 配置
    ├── handlers/        # 处理器
    ├── middleware/      # 中间件
    ├── models/          # 数据模型
    └── routes/          # 路由
```

## 许可证

MIT License