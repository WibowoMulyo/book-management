package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func FormatValidationErrors(err error) []string {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			errors = append(errors, formatFieldError(fieldError))
		}
	}

	return errors
}

func formatFieldError(fieldError validator.FieldError) string {
	field := strings.ToLower(fieldError.Field())
	tag := fieldError.Tag()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters/value", field, fieldError.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters/value", field, fieldError.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
