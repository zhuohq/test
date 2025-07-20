package validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// CustomValidator 自定义验证器
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator 创建新的自定义验证器
func NewCustomValidator() *CustomValidator {
	v := validator.New()
	
	// 注册自定义验证规则
	v.RegisterValidation("password", validatePassword)
	v.RegisterValidation("phone", validatePhone)
	v.RegisterValidation("username", validateUsername)
	
	return &CustomValidator{validator: v}
}

// Validate 验证结构体
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// validatePassword 密码验证
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	
	// 密码至少8位，包含大小写字母和数字
	if len(password) < 8 {
		return false
	}
	
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	
	return hasUpper && hasLower && hasNumber
}

// validatePhone 手机号验证
func validatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	
	// 中国手机号验证
	pattern := `^1[3-9]\d{9}$`
	matched, _ := regexp.MatchString(pattern, phone)
	
	return matched
}

// validateUsername 用户名验证
func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	
	// 用户名3-20位，只能包含字母、数字、下划线
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	
	pattern := `^[a-zA-Z0-9_]+$`
	matched, _ := regexp.MatchString(pattern, username)
	
	return matched
}

// ValidationError 验证错误结构
type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

// GetValidationErrors 获取验证错误详情
func GetValidationErrors(err error) []ValidationError {
	var errors []ValidationError
	
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := e.Field()
			tag := e.Tag()
			value := e.Value()
			
			// 将字段名转换为小写（JSON格式）
			field = strings.ToLower(field[:1]) + field[1:]
			
			message := getErrorMessage(field, tag)
			
			errors = append(errors, ValidationError{
				Field:   field,
				Tag:     tag,
				Value:   reflect.ValueOf(value).String(),
				Message: message,
			})
		}
	}
	
	return errors
}

// getErrorMessage 获取错误消息
func getErrorMessage(field, tag string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		return field + " must be at least " + tag + " characters"
	case "max":
		return field + " must be at most " + tag + " characters"
	case "password":
		return field + " must be at least 8 characters with uppercase, lowercase and number"
	case "phone":
		return field + " must be a valid phone number"
	case "username":
		return field + " must be 3-20 characters with only letters, numbers and underscore"
	default:
		return field + " is invalid"
	}
}