package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-restful/helper"
)

type ValidationError struct {
	Message string
	Field   string
}

type ErrorBag struct {
	Errors []ValidationError
}

func (e *ErrorBag) addError(err *ValidationError) {
	e.Errors = append(e.Errors, *err)
}

func (e *ErrorBag) Map() map[string][]string {
	var errors = make(map[string][]string)

	for _, err := range e.Errors {
		field := helper.ToSnakeCase(err.Field)
		if _, ok := errors[field]; ok {
			errors[field] = append(errors[field], err.Message)
		} else {
			errors[field] = []string{err.Message}
		}
	}

	return errors
}

func newErrorBag() *ErrorBag {
	return &ErrorBag{}
}

func ParseErrors(err *validator.ValidationErrors) *ErrorBag {
	var errorBag = newErrorBag()

	for _, e := range *err {
		field := helper.ToSnakeCase(e.Field())
		switch e.Tag() {
		case "required":
			errorBag.addError(&ValidationError{Field: field, Message: "The " + field + " field is required."})
		case "email":
			errorBag.addError(&ValidationError{Field: field, Message: "The " + field + " field must be a valid email address."})
		case "max":
			errorBag.addError(&ValidationError{Field: field, Message: "The " + field + " must have max length of " + e.Param()})
		case "min":
			errorBag.addError(&ValidationError{Field: field, Message: "The " + field + " must have min length of " + e.Param()})
		case "eqfield":
			errorBag.addError(&ValidationError{Field: field, Message: "The " + field + " must be same with " + helper.ToSnakeCase(e.Param())})
		default:
			errorBag.addError(&ValidationError{Field: field, Message: "The " + field + " field is invalid."})
		}
	}

	return errorBag
}
