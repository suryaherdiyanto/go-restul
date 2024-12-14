package response

import "github.com/go-restful/app/model"

type UserDataResponse struct {
	FirstName string      `json:"first_name"`
	LastName  interface{} `json:"last_name"`
	Email     string      `json:"email"`
}

func NewUserDataResponse(user *model.User) UserDataResponse {
	return UserDataResponse{
		FirstName: user.FirstName,
		Email:     user.Email,
		LastName:  user.LastName.String,
	}
}

func NewUsersDataResponse(users *[]model.User) []UserDataResponse {
	var usersResponse []UserDataResponse
	for _, user := range *users {
		usersResponse = append(usersResponse, NewUserDataResponse(&user))
	}

	return usersResponse
}
