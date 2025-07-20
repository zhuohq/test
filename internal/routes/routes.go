package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go-echo-app/internal/handlers"
)

// SetupRoutes 设置所有路由
func SetupRoutes(e *echo.Echo) {
	// 健康检查端点
	e.GET("/health", healthCheck)

	// API路由组
	api := e.Group("/api/v1")
	
	// 用户相关路由
	api.GET("/users", handlers.GetUsers)
	api.GET("/users/:id", handlers.GetUser)
	api.POST("/users", handlers.CreateUser)
	api.PUT("/users/:id", handlers.UpdateUser)
	api.DELETE("/users/:id", handlers.DeleteUser)

	// 根路径
	e.GET("/", homePage)
}

// healthCheck 健康检查处理函数
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "Server is running",
		"version": "1.0.0",
	})
}

// homePage 首页处理函数
func homePage(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Welcome to Go Echo API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"health": "/health",
			"users":  "/api/v1/users",
		},
	})
}