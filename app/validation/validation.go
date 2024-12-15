package validation

import "github.com/go-playground/validator/v10"

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
		if _, ok := errors[err.Field]; ok {
			errors[err.Field] = append(errors[err.Field], err.Message)
		} else {
			errors[err.Field] = []string{err.Message}
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
		switch e.Tag() {
		case "required":
			errorBag.addError(&ValidationError{Field: e.Field(), Message: "The " + e.Field() + " field is required."})
		case "email":
			errorBag.addError(&ValidationError{Field: e.Field(), Message: "The " + e.Field() + " field must be a valid email address."})
		default:
			errorBag.addError(&ValidationError{Field: e.Field(), Message: "The " + e.Field() + " field is invalid."})
		}
	}

	return errorBag
}
