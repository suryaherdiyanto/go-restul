package response

import (
	"net/http"
)

type NotfoundResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewNotfoundResponse(status string, message string) NotfoundResponse {
	return NotfoundResponse{
		Status:  status,
		Message: message,
	}
}

func HandleNotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	res := NewNotfoundResponse(http.StatusText(http.StatusNotFound), message)
	JsonResponse(w, res)
}
