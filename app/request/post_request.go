package request

import "io"

type PostRequest struct {
	Title    string `json:"title" validate:"required,max=100"`
	Category string `json:"category" validate:"required,max=50"`
	Content  string `json:"content" validate:"required"`
	UserID   int    `json:"user_id" validate:"required,number,min=1"`
}

type PostUpdateRequest struct {
	Title    string `json:"title" validate:"required,max=100"`
	Category string `json:"category" validate:"required,max=50"`
	Content  string `json:"content" validate:"required"`
	UserID   int    `json:"user_id" validate:"required,number,min=1"`
}

func NewPostRequest(body io.Reader) (*PostRequest, error) {
	post := &PostRequest{}
	err := Parse(body, post)

	return post, err
}

func NewPostUpdateRequest(body io.Reader) (*PostUpdateRequest, error) {
	post := &PostUpdateRequest{}
	err := Parse(body, post)

	return post, err
}
