package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

func (ve ValidationErrors) Error() string {
	var sb strings.Builder
	for _, err := range ve.Errors {
		sb.WriteString(fmt.Sprintf("Field: %s, Error: %s\n", err.Field, err.Message))
	}
	return sb.String()
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var errors []ValidationError
		t := reflect.TypeOf(i).Elem()
		for _, err := range err.(validator.ValidationErrors) {
			field, ok := t.FieldByName(err.StructField())
			fieldName := err.StructField()
			if ok {
				if field.Tag.Get("param") != "" {
					fieldName = field.Tag.Get("param")
				} else if field.Tag.Get("query") != "" {
					fieldName = field.Tag.Get("query")
				}
			}

			errors = append(errors, ValidationError{
				Field:   fieldName,
				Message: cv.formatErrorMessage(fieldName, err),
			})
		}
		return ValidationErrors{Errors: errors}
	}

	return nil
}

// Custom error message formatting
func (cv *CustomValidator) formatErrorMessage(fieldName string, err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "The " + fieldName + " field is required"
	case "email":
		return "The " + fieldName + " field must be a valid email"
	case "gte":
		return "The " + fieldName + " field must be greater than or equal to " + err.Param()
	case "lte":
		return "The " + fieldName + " field must be less than or equal to " + err.Param()
	case "min":
		return "The " + fieldName + " field must be at least " + err.Param() + " characters long"
	case "max":
		return "The " + fieldName + " field must be at most " + err.Param() + " characters long"
	case "len":
		return "The " + fieldName + " field must be exactly " + err.Param() + " characters long"
	case "alphanum":
		return "The " + fieldName + " field must contain only alphanumeric characters"
	case "contains":
		return "The " + fieldName + " field must contain " + err.Param()
	default:
		return "The " + fieldName + " field is invalid"
	}
}

func BindAndValidate(c echo.Context, payload interface{}) error {
	if err := c.Bind(payload); err != nil {
		return err
	}

	if err := c.Validate(payload); err != nil {
		return err
	}

	return nil
}
