package response

import "net/http"

type SuccessResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (r *SuccessResponse) GetStatus() int {
	return r.Status
}

func (r *SuccessResponse) GetMessage() string {
	return ""
}

func (r *SuccessResponse) GetData() interface{} {
	return r.Data
}

func NewSuccessResponse(data interface{}) Response {
	return &SuccessResponse{Status: http.StatusOK, Data: data}
}
