package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// User 用户结构体
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUsers 获取所有用户
func GetUsers(c echo.Context) error {
	users := []User{
		{ID: "1", Name: "John Doe", Email: "john@example.com"},
		{ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
		{ID: "3", Name: "Bob Johnson", Email: "bob@example.com"},
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
		"count": len(users),
	})
}

// GetUser 根据ID获取用户
func GetUser(c echo.Context) error {
	id := c.Param("id")
	
	// 模拟从数据库获取用户
	user := User{
		ID:    id,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	
	return c.JSON(http.StatusOK, user)
}

// CreateUser 创建新用户
func CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	
	// 这里应该添加用户到数据库
	// 模拟设置ID
	user.ID = "new-id"
	
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	})
}

// UpdateUser 更新用户信息
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(User)
	
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	
	user.ID = id
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    user,
	})
}

// DeleteUser 删除用户
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	
	// 这里应该从数据库删除用户
	
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
		"id":      id,
	})
}