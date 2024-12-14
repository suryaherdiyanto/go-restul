package response

import (
	"net/http"
)

type NotfoundResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (r *NotfoundResponse) GetStatus() int {
	return r.Status
}

func (r *NotfoundResponse) GetMessage() string {
	return r.Message
}

func (r *NotfoundResponse) GetData() interface{} {
	return nil
}

func NewNotfoundResponse(message string) Response {
	return &NotfoundResponse{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func HandleNotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	res := NewNotfoundResponse(message)
	JsonResponse(w, res)
}
