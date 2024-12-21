package request

import "io"

type PostRequest struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	UserID   int    `json:"user_id"`
}

type PostUpdateRequest struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	UserID   int    `json:"user_id"`
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
