package response

import "net/http"

type InternalServerError struct {
	Status    int
	Message   string
	Exception interface{}
}

func (e *InternalServerError) GetStatus() int {
	return e.Status
}

func (e *InternalServerError) GetMessage() string {
	return e.Message
}

func (e *InternalServerError) GetData() interface{} {
	return e.Exception
}

func NewInternalServerError(message string, exception interface{}) Response {
	return &InternalServerError{
		Status:    http.StatusInternalServerError,
		Message:   message,
		Exception: exception,
	}
}
