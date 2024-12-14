package request

type UserRequest struct {
	FirstName string      `json:"first_name"`
	LastName  interface{} `json:"last_name"`
	Email     string      `json:"email"`
}
