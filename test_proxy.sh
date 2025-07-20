#!/bin/bash

# HTTP中转API测试脚本

echo "=== HTTP中转API测试 ==="

# 启动服务器（如果还没有启动）
echo "启动服务器..."
go run main.go &
SERVER_PID=$!

# 等待服务器启动
sleep 3

echo ""
echo "1. 测试基本GET请求中转"
echo "目标: https://httpbin.org/get"
curl -X GET "http://localhost:8080/api/v1/proxy?target=https://httpbin.org/get" \
  -H "Content-Type: application/json"

echo ""
echo ""
echo "2. 测试POST请求中转"
echo "目标: https://httpbin.org/post"
curl -X POST "http://localhost:8080/api/v1/proxy?target=https://httpbin.org/post" \
  -H "Content-Type: application/json" \
  -d '{"name": "test", "message": "hello world"}'

echo ""
echo ""
echo "3. 测试带自定义请求头的GET请求"
echo "目标: https://httpbin.org/headers"
curl -X GET "http://localhost:8080/api/v1/proxy?target=https://httpbin.org/headers" \
  -H "X-Custom-Header: test-value" \
  -H "User-Agent: Proxy-Test/1.0"

echo ""
echo ""
echo "4. 测试带配置的POST请求"
echo "目标: https://httpbin.org/post"
curl -X POST "http://localhost:8080/api/v1/proxy/config" \
  -H "Content-Type: application/json" \
  -d '{
    "target_url": "https://httpbin.org/post",
    "method": "POST",
    "headers": {
      "X-Custom-Header": "config-test",
      "Authorization": "Bearer test-token"
    },
    "body": {
      "test": "data",
      "nested": {
        "key": "value"
      }
    },
    "timeout": 10
  }'

echo ""
echo ""
echo "5. 测试错误情况 - 缺少目标URL"
curl -X GET "http://localhost:8080/api/v1/proxy" \
  -H "Content-Type: application/json"

echo ""
echo ""
echo "6. 测试错误情况 - 无效的URL"
curl -X GET "http://localhost:8080/api/v1/proxy?target=invalid-url" \
  -H "Content-Type: application/json"

echo ""
echo ""
echo "=== 测试完成 ==="

# 停止服务器
echo "停止服务器..."
kill $SERVER_PID 2>/dev/null