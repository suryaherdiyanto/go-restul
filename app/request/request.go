package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-restful/app/validation"
)

func Validate(data interface{}) (*validation.ErrorBag, bool) {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := validation.ParseErrors(&validationErrors)
		return errors, false
	}

	return &validation.ErrorBag{}, true
}
