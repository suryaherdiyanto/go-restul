package response

import "net/http"

type UnAuthorizedResponse struct {
	Status  int
	Message string
}

func (r *UnAuthorizedResponse) GetStatus() int {
	return r.Status
}

func (r *UnAuthorizedResponse) GetMessage() string {
	return r.Message
}

func (r *UnAuthorizedResponse) GetData() interface{} {
	return nil
}

func NewUnAuthorizedResponse(message string) *UnAuthorizedResponse {
	return &UnAuthorizedResponse{
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}
