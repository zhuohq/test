package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response 统一响应结构
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse 成功响应
func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse 错误响应
func ErrorResponse(c echo.Context, statusCode int, message string, err string) error {
	return c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// BadRequest 400错误响应
func BadRequest(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusBadRequest, message, "Bad Request")
}

// Unauthorized 401错误响应
func Unauthorized(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusUnauthorized, message, "Unauthorized")
}

// Forbidden 403错误响应
func Forbidden(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusForbidden, message, "Forbidden")
}

// NotFound 404错误响应
func NotFound(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusNotFound, message, "Not Found")
}

// InternalServerError 500错误响应
func InternalServerError(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusInternalServerError, message, "Internal Server Error")
}

// ValidationError 验证错误响应
func ValidationError(c echo.Context, message string) error {
	return ErrorResponse(c, http.StatusUnprocessableEntity, message, "Validation Error")
}