package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CustomMiddleware 自定义中间件配置
func CustomMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		// 日志中间件
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency}\n",
		}),
		
		// 恢复中间件
		middleware.Recover(),
		
		// CORS中间件
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}),
		
		// 请求ID中间件
		middleware.RequestID(),
		
		// 超时中间件
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: 30 * time.Second,
		}),
	}
}

// AuthMiddleware 认证中间件
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 这里可以添加JWT验证逻辑
			// token := c.Request().Header.Get("Authorization")
			// if token == "" {
			//     return c.JSON(401, map[string]string{"error": "Unauthorized"})
			// }
			
			return next(c)
		}
	}
}

// RateLimitMiddleware 限流中间件
func RateLimitMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 这里可以添加限流逻辑
			// 例如使用令牌桶算法
			
			return next(c)
		}
	}
}