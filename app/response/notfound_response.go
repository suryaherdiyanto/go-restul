package response

import (
	"encoding/json"
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

func HandlerNotFound(w http.ResponseWriter, enc *json.Encoder, message string) {
	w.WriteHeader(http.StatusNotFound)
	res := NewNotfoundResponse(http.StatusText(http.StatusNotFound), message)
	enc.Encode(res)
}
