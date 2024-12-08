package response

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
