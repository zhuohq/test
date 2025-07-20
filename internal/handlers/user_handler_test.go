package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// 创建Echo实例
	e := echo.New()
	
	// 创建请求
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// 执行处理函数
	err := GetUsers(c)
	
	// 断言
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	
	// 验证响应包含用户数据
	responseBody := rec.Body.String()
	assert.Contains(t, responseBody, "users")
	assert.Contains(t, responseBody, "count")
}

func TestGetUser(t *testing.T) {
	// 创建Echo实例
	e := echo.New()
	
	// 创建请求
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	// 执行处理函数
	err := GetUser(c)
	
	// 断言
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	
	// 验证响应包含用户ID
	responseBody := rec.Body.String()
	assert.Contains(t, responseBody, "1")
}

func TestCreateUser(t *testing.T) {
	// 创建Echo实例
	e := echo.New()
	
	// 创建请求数据
	userData := `{"name":"Test User","email":"test@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// 执行处理函数
	err := CreateUser(c)
	
	// 断言
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	
	// 验证响应包含成功消息
	responseBody := rec.Body.String()
	assert.Contains(t, responseBody, "User created successfully")
}

func TestUpdateUser(t *testing.T) {
	// 创建Echo实例
	e := echo.New()
	
	// 创建请求数据
	userData := `{"name":"Updated User","email":"updated@example.com"}`
	req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(userData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	// 执行处理函数
	err := UpdateUser(c)
	
	// 断言
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	
	// 验证响应包含成功消息
	responseBody := rec.Body.String()
	assert.Contains(t, responseBody, "User updated successfully")
}

func TestDeleteUser(t *testing.T) {
	// 创建Echo实例
	e := echo.New()
	
	// 创建请求
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	
	// 执行处理函数
	err := DeleteUser(c)
	
	// 断言
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	
	// 验证响应包含成功消息
	responseBody := rec.Body.String()
	assert.Contains(t, responseBody, "User deleted successfully")
}