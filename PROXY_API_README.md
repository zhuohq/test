# HTTP中转API使用说明

这个项目提供了一个HTTP中转API，可以将HTTP请求转发到目标服务器并返回响应。

## 功能特性

- 支持所有HTTP方法（GET, POST, PUT, DELETE, PATCH等）
- 自动转发请求头和请求体
- 支持自定义超时设置
- 提供两种使用方式：简单模式和配置模式
- 错误处理和响应状态码保持

## API端点

### 1. 简单中转API

**端点**: `/api/v1/proxy`

**支持方法**: GET, POST, PUT, DELETE, PATCH

**使用方式**:
- 通过查询参数指定目标URL: `?target=https://example.com/api`
- 通过请求头指定目标URL: `X-Target-URL: https://example.com/api`

#### 示例

```bash
# GET请求
curl -X GET "http://localhost:8080/api/v1/proxy?target=https://httpbin.org/get"

# POST请求
curl -X POST "http://localhost:8080/api/v1/proxy?target=https://httpbin.org/post" \
  -H "Content-Type: application/json" \
  -d '{"name": "test", "message": "hello"}'

# 使用请求头指定目标URL
curl -X GET "http://localhost:8080/api/v1/proxy" \
  -H "X-Target-URL: https://httpbin.org/headers" \
  -H "X-Custom-Header: test-value"
```

### 2. 配置模式中转API

**端点**: `/api/v1/proxy/config`

**支持方法**: POST

**请求体格式**:
```json
{
  "target_url": "https://example.com/api",
  "method": "POST",
  "headers": {
    "Authorization": "Bearer token",
    "X-Custom-Header": "value"
  },
  "body": {
    "key": "value"
  },
  "timeout": 30
}
```

#### 字段说明

- `target_url` (必需): 目标服务器URL
- `method` (可选): HTTP方法，默认为当前请求的方法
- `headers` (可选): 自定义请求头
- `body` (可选): 请求体内容
- `timeout` (可选): 超时时间（秒），默认30秒

#### 示例

```bash
curl -X POST "http://localhost:8080/api/v1/proxy/config" \
  -H "Content-Type: application/json" \
  -d '{
    "target_url": "https://httpbin.org/post",
    "method": "POST",
    "headers": {
      "Authorization": "Bearer test-token",
      "X-Custom-Header": "config-test"
    },
    "body": {
      "test": "data",
      "nested": {
        "key": "value"
      }
    },
    "timeout": 10
  }'
```

## 响应格式

### 成功响应

中转API会保持原始响应的状态码、响应头和响应体。

### 错误响应

```json
{
  "error": "错误描述"
}
```

#### 常见错误

- `400 Bad Request`: 缺少目标URL或URL格式无效
- `502 Bad Gateway`: 目标服务器无响应或连接失败
- `500 Internal Server Error`: 服务器内部错误

## 使用场景

1. **API网关**: 作为微服务架构中的API网关
2. **CORS代理**: 解决跨域请求问题
3. **请求转发**: 将请求转发到不同的后端服务
4. **负载均衡**: 配合负载均衡器使用
5. **调试工具**: 用于调试和测试API

## 安全注意事项

1. **URL验证**: 确保目标URL是可信的
2. **请求头过滤**: 某些敏感请求头不会被转发
3. **超时设置**: 避免长时间等待响应
4. **错误处理**: 妥善处理各种错误情况

## 测试

运行测试脚本：

```bash
chmod +x test_proxy.sh
./test_proxy.sh
```

## 启动服务器

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

## 项目结构

```
.
├── main.go                    # 主程序入口
├── internal/
│   └── handlers/
│       ├── proxy_handler.go   # HTTP中转处理器
│       └── user_handler.go    # 用户相关处理器
├── test_proxy.sh             # 测试脚本
└── PROXY_API_README.md       # 本文档
```