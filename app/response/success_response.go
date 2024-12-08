package response

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewSuccessResponse(status string, data interface{}) SuccessResponse {
	return SuccessResponse{Status: status, Data: data}
}
