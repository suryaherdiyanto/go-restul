package request

import (
	"io"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginRequest(body io.Reader) (*LoginRequest, error) {
	post := &LoginRequest{}
	err := Parse(body, post)

	return post, err
}
