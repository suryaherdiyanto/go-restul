package request

import (
	"encoding/json"
	"io"

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

func Parse(r io.Reader, data interface{}) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(data)

	if err != nil {
		return err
	}

	return nil
}
