# HTTP中转API使用示例

## 快速开始

### 1. 启动服务器

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

### 2. 基本使用示例

#### 示例1: 转发GET请求到httpbin.org

```bash
curl -X GET "http://localhost:8080/api/v1/proxy?target=https://httpbin.org/get"
```

**预期响应**:
```json
{
  "args": {},
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/1.1"
  },
  "origin": "127.0.0.1",
  "url": "https://httpbin.org/get"
}
```

#### 示例2: 转发POST请求

```bash
curl -X POST "http://localhost:8080/api/v1/proxy?target=https://httpbin.org/post" \
  -H "Content-Type: application/json" \
  -d '{"name": "张三", "age": 25, "city": "北京"}'
```

#### 示例3: 使用请求头指定目标URL

```bash
curl -X GET "http://localhost:8080/api/v1/proxy" \
  -H "X-Target-URL: https://httpbin.org/headers" \
  -H "X-Custom-Header: test-value" \
  -H "Authorization: Bearer my-token"
```

### 3. 配置模式示例

#### 示例4: 使用配置模式转发请求

```bash
curl -X POST "http://localhost:8080/api/v1/proxy/config" \
  -H "Content-Type: application/json" \
  -d '{
    "target_url": "https://httpbin.org/post",
    "method": "POST",
    "headers": {
      "Authorization": "Bearer your-token-here",
      "X-Custom-Header": "proxy-test",
      "Content-Type": "application/json"
    },
    "body": {
      "user": {
        "id": 123,
        "name": "李四",
        "email": "lisi@example.com"
      },
      "action": "create",
      "timestamp": "2024-01-15T10:30:00Z"
    },
    "timeout": 15
  }'
```

### 4. 实际应用场景示例

#### 场景1: API网关

假设你有多个微服务，可以通过中转API统一入口：

```bash
# 转发到用户服务
curl -X GET "http://localhost:8080/api/v1/proxy?target=http://user-service:8081/users"

# 转发到订单服务
curl -X POST "http://localhost:8080/api/v1/proxy?target=http://order-service:8082/orders" \
  -H "Content-Type: application/json" \
  -d '{"product_id": 123, "quantity": 2}'

# 转发到支付服务
curl -X POST "http://localhost:8080/api/v1/proxy?target=http://payment-service:8083/payments" \
  -H "Content-Type: application/json" \
  -d '{"amount": 99.99, "currency": "CNY"}'
```

#### 场景2: CORS代理

解决前端跨域问题：

```bash
# 前端JavaScript代码
fetch('http://localhost:8080/api/v1/proxy?target=https://api.example.com/data', {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json'
  }
})
.then(response => response.json())
.then(data => console.log(data));
```

#### 场景3: 调试外部API

```bash
# 测试GitHub API
curl -X GET "http://localhost:8080/api/v1/proxy?target=https://api.github.com/users/octocat"

# 测试JSONPlaceholder API
curl -X POST "http://localhost:8080/api/v1/proxy?target=https://jsonplaceholder.typicode.com/posts" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "测试文章",
    "body": "这是一篇测试文章的内容",
    "userId": 1
  }'
```

### 5. 错误处理示例

#### 错误1: 缺少目标URL

```bash
curl -X GET "http://localhost:8080/api/v1/proxy"
```

**响应**:
```json
{
  "error": "Missing target URL. Please provide 'target' query parameter or 'X-Target-URL' header"
}
```

#### 错误2: 无效的URL

```bash
curl -X GET "http://localhost:8080/api/v1/proxy?target=invalid-url"
```

**响应**:
```json
{
  "error": "Invalid target URL"
}
```

#### 错误3: 目标服务器无响应

```bash
curl -X GET "http://localhost:8080/api/v1/proxy?target=https://nonexistent-domain-12345.com"
```

**响应**:
```json
{
  "error": "Failed to forward request: Get \"https://nonexistent-domain-12345.com\": dial tcp: lookup nonexistent-domain-12345.com: no such host"
}
```

### 6. 高级用法

#### 使用环境变量

```bash
# 设置目标URL环境变量
export TARGET_API="https://api.example.com"

# 使用环境变量
curl -X GET "http://localhost:8080/api/v1/proxy?target=$TARGET_API/users"
```

#### 批量请求示例

```bash
#!/bin/bash

# 批量测试多个API
apis=(
  "https://httpbin.org/get"
  "https://httpbin.org/post"
  "https://jsonplaceholder.typicode.com/posts/1"
  "https://api.github.com/users/octocat"
)

for api in "${apis[@]}"; do
  echo "Testing: $api"
  curl -s -X GET "http://localhost:8080/api/v1/proxy?target=$api" | jq '.'
  echo "---"
done
```

### 7. 性能测试

```bash
# 使用ab进行性能测试
ab -n 100 -c 10 "http://localhost:8080/api/v1/proxy?target=https://httpbin.org/get"

# 使用wrk进行压力测试
wrk -t12 -c400 -d30s "http://localhost:8080/api/v1/proxy?target=https://httpbin.org/get"
```

这些示例展示了HTTP中转API的各种使用场景和功能。你可以根据实际需求调整和扩展这些示例。