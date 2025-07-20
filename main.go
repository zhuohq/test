package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"go-echo-app/internal/config"
	"go-echo-app/internal/middleware"
	"go-echo-app/internal/routes"
	"go-echo-app/pkg/validator"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 创建Echo实例
	e := echo.New()

	// 设置自定义验证器
	e.Validator = validator.NewCustomValidator()

	// 添加中间件
	for _, m := range middleware.CustomMiddleware() {
		e.Use(m)
	}

	// 设置路由
	routes.SetupRoutes(e)

	// 启动服务器
	serverAddr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", serverAddr)
	log.Fatal(e.Start(serverAddr))
}