package request

import (
	"io"
)

type UserRequest struct {
	FirstName            string      `json:"first_name" validate:"required"`
	LastName             interface{} `json:"last_name"`
	Password             string      `json:"password" validate:"required,min=8,eqfield=PasswordConfirmation"`
	PasswordConfirmation string      `json:"password_confirmation" validate:"required,min=8"`
	Email                string      `json:"email" validate:"required,email"`
}

type UserUpdateRequest struct {
	FirstName string `json:"first_name" validate:"required,max=50"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email,max=50"`
}

func NewUserRequest(body io.Reader) (*UserRequest, error) {
	r := &UserRequest{}
	Parse(body, r)

	return r, nil
}

func NewUserUpdateRequest(body io.Reader) (*UserUpdateRequest, error) {
	r := &UserUpdateRequest{}
	Parse(body, r)

	return r, nil
}
