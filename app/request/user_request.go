package request

type UserRequest struct {
	FirstName string      `json:"first_name" validate:"required"`
	LastName  interface{} `json:"last_name"`
	Email     string      `json:"email" validate:"required,email"`
}
