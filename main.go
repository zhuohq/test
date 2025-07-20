package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 创建Echo实例
	e := echo.New()

	// 添加中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 设置路由
	setupRoutes(e)

	// 启动服务器
	log.Fatal(e.Start(":8080"))
}

func setupRoutes(e *echo.Echo) {
	// 健康检查端点
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
			"message": "Server is running",
		})
	})

	// API路由组
	api := e.Group("/api/v1")
	
	// 用户相关路由
	api.GET("/users", getUsers)
	api.GET("/users/:id", getUser)
	api.POST("/users", createUser)
	api.PUT("/users/:id", updateUser)
	api.DELETE("/users/:id", deleteUser)

	// 根路径
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome to Go Echo API",
			"version": "1.0.0",
		})
	})
}

// 用户相关的处理函数
func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": []map[string]string{
			{"id": "1", "name": "John Doe", "email": "john@example.com"},
			{"id": "2", "name": "Jane Smith", "email": "jane@example.com"},
		},
	})
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"id":    id,
		"name":  "John Doe",
		"email": "john@example.com",
	})
}

func createUser(c echo.Context) error {
	user := make(map[string]interface{})
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	})
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	user := make(map[string]interface{})
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"id":      id,
		"user":    user,
	})
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
		"id":      id,
	})
}