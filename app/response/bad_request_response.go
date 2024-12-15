package response

import "net/http"

type BadRequestResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Errors  interface{} `json:"errors"`
}

func (r *BadRequestResponse) GetMessage() string {
	return r.Message
}

func (r *BadRequestResponse) GetStatus() int {
	return r.Status
}

func (r *BadRequestResponse) GetData() interface{} {
	return r.Errors
}

func NewBadRequestResponse(message string, errors interface{}) Response {
	return &BadRequestResponse{
		Message: message,
		Status:  http.StatusBadRequest,
		Errors:  errors,
	}
}
