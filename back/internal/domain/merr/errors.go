package merr

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ResponseError struct {
	StatusCode int            `json:"-"`
	Error      error          `json:"-"`
	Message    string         `json:"message"`
	Info       map[string]any `json:"info,omitempty"`
}

func NewResponseError(status int, err error) *ResponseError {
	return &ResponseError{
		StatusCode: status,
		Message:    err.Error(),
		Error:      err,
		Info:       nil,
	}
}

func NewReponseErrorInfo(status int, err error, info map[string]any) *ResponseError {
	return &ResponseError{
		StatusCode: status,
		Message:    err.Error(),
		Info:       info,
	}
}

func BindValidationErrorsToResponse(err error) *ResponseError {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		return &ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    "Validation failed",
			Info:       mapValidationErrors(validationErrs),
		}
	}

	return &ResponseError{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}
}

func mapValidationErrors(errs validator.ValidationErrors) map[string]any {
	errorMap := make(map[string]any)

	for _, err := range errs {
		var fieldName string = strings.ToLower(err.Field())
		var tag string = err.Tag()

		if errorMap[fieldName] == nil {
			errorMap[fieldName] = make(map[string]string)
		}

		fieldErrors := errorMap[fieldName].(map[string]string)
		fieldErrors[tag] = getErrorMessage(err)
	}

	return errorMap
}

func getErrorMessage(err validator.FieldError) string {
	var field string = err.Field()
	var tag string = err.Tag()
	var param string = err.Param()

	var lowerField string = strings.ToLower(field)

	switch tag {
	case "required":
		return fmt.Sprintf("The %s is required", lowerField)
	case "email":
		return fmt.Sprintf("The %s must be a valid email address", lowerField)
	case "min":
		return fmt.Sprintf("The %s must be at least %s characters", lowerField, param)
	case "max":
		return fmt.Sprintf("The %s must not exceed %s characters", lowerField, param)
	case "len":
		return fmt.Sprintf("The %s must be exactly %s characters", lowerField, param)
	case "numeric":
		return fmt.Sprintf("The %s must contain only numbers", lowerField)
	case "oneof":
		return fmt.Sprintf("The %s must be one of: %s", lowerField, param)
	case "eqfield":
		return fmt.Sprintf("The %s must match %s", lowerField, strings.ToLower(param))
	case "containsuppercase":
		return fmt.Sprintf("The %s must contain at least one uppercase letter", lowerField)
	case "containslowercase":
		return fmt.Sprintf("The %s must contain at least one lowercase letter", lowerField)
	case "containsdigit":
		return fmt.Sprintf("The %s must contain at least one number", lowerField)
	case "containsspecial":
		return fmt.Sprintf("The %s must contain at least one special character", lowerField)
	default:
		return fmt.Sprintf("The %s is invalid", lowerField)
	}
}
