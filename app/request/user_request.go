package request

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-restful/app/validation"
)

type JsonRequest interface {
	Validate() (*validation.ErrorBag, bool)
}

type UserRequest struct {
	FirstName string      `json:"first_name" validate:"required"`
	LastName  interface{} `json:"last_name"`
	Email     string      `json:"email" validate:"required,email"`
}

type UserUpdateRequest struct {
	FirstName string `json:"first_name" validate:"required,max=50"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email,max=50"`
}

func NewUserRequest(body io.Reader) (*UserRequest, error) {
	r := &UserRequest{}
	dec := json.NewDecoder(body)
	err := dec.Decode(r)

	if err != nil {
		return r, err
	}

	return r, nil
}

func NewUserUpdateRequest(body io.Reader) (*UserUpdateRequest, error) {
	r := &UserUpdateRequest{}
	dec := json.NewDecoder(body)
	err := dec.Decode(r)

	if err != nil {
		return r, err
	}

	fmt.Printf("%v", r)

	return r, nil
}
