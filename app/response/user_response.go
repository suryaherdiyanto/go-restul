package response

import "github.com/go-restful/app/model"

type UserResponse struct {
	FirstName string      `json:"first_name"`
	LastName  interface{} `json:"last_name"`
	Email     string      `json:"email"`
}

func NewUserResponse(user model.User) UserResponse {
	return UserResponse{
		FirstName: user.FirstName,
		Email:     user.Email,
		LastName:  user.LastName.String,
	}
}

func NewUsersResponse(users []model.User) []UserResponse {
	var usersResponse []UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, NewUserResponse(user))
	}

	return usersResponse
}
