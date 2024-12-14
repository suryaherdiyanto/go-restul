package response

import "net/http"

type CreatedResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *CreatedResponse) GetMessage() string {
	return r.Message
}

func (r *CreatedResponse) GetStatus() int {
	return r.Status
}

func (r *CreatedResponse) GetData() interface{} {
	return r.Data
}

func NewCreatedResponse(message string, data interface{}) Response {
	return &CreatedResponse{Status: http.StatusCreated, Message: message, Data: data}
}
