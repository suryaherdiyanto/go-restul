package request

import (
	"io"
)

type CategoryRequest struct {
}

type CategoryUpdateRequest struct {
}

func NewCategoryRequest(body io.Reader) (*CategoryRequest, error) {
	r := &CategoryRequest{}
	Parse(body, r)

	return r, nil
}

func NewCategoryUpdateRequest(body io.Reader) (*CategoryUpdateRequest, error) {
	r := &CategoryUpdateRequest{}
	Parse(body, r)

	return r, nil
}
