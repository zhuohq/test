package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo/v4"
)

// ProxyRequest 处理HTTP中转请求
func ProxyRequest(c echo.Context) error {
	// 从请求头或查询参数中获取目标URL
	targetURL := c.QueryParam("target")
	if targetURL == "" {
		targetURL = c.Request().Header.Get("X-Target-URL")
	}
	
	if targetURL == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Missing target URL. Please provide 'target' query parameter or 'X-Target-URL' header",
		})
	}

	// 解析目标URL
	_, err := url.Parse(targetURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid target URL",
		})
	}

	// 读取请求体
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to read request body",
		})
	}

	// 创建转发请求
	req, err := http.NewRequest(c.Request().Method, targetURL, bytes.NewBuffer(body))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create request",
		})
	}

	// 复制请求头（排除一些不应该转发的头）
	for key, values := range c.Request().Header {
		if key != "Host" && key != "X-Target-URL" {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}

	// 设置超时
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{
			"error": "Failed to forward request: " + err.Error(),
		})
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to read response body",
		})
	}

	// 复制响应头
	for key, values := range resp.Header {
		for _, value := range values {
			c.Response().Header().Add(key, value)
		}
	}

	// 返回响应
	return c.Blob(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}

// ProxyRequestWithConfig 带配置的HTTP中转请求
func ProxyRequestWithConfig(c echo.Context) error {
	// 从请求体获取配置
	var config struct {
		TargetURL string            `json:"target_url"`
		Method    string            `json:"method,omitempty"`
		Headers   map[string]string `json:"headers,omitempty"`
		Body      interface{}       `json:"body,omitempty"`
		Timeout   int               `json:"timeout,omitempty"`
	}

	if err := c.Bind(&config); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if config.TargetURL == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Missing target_url in request body",
		})
	}

	// 解析目标URL
	_, err := url.Parse(config.TargetURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid target URL",
		})
	}

	// 确定HTTP方法
	method := config.Method
	if method == "" {
		method = c.Request().Method
	}

	// 准备请求体
	var body io.Reader
	if config.Body != nil {
		// 这里可以根据需要将config.Body转换为JSON或其他格式
		bodyBytes, _ := json.Marshal(config.Body)
		body = bytes.NewBuffer(bodyBytes)
	}

	// 创建转发请求
	req, err := http.NewRequest(method, config.TargetURL, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create request",
		})
	}

	// 设置自定义请求头
	for key, value := range config.Headers {
		req.Header.Set(key, value)
	}

	// 设置默认Content-Type（如果没有提供）
	if req.Header.Get("Content-Type") == "" && body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 设置超时
	timeout := 30 * time.Second
	if config.Timeout > 0 {
		timeout = time.Duration(config.Timeout) * time.Second
	}

	client := &http.Client{
		Timeout: timeout,
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{
			"error": "Failed to forward request: " + err.Error(),
		})
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to read response body",
		})
	}

	// 复制响应头
	for key, values := range resp.Header {
		for _, value := range values {
			c.Response().Header().Add(key, value)
		}
	}

	// 返回响应
	return c.Blob(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}