package resource

import (
	"github.com/go-restful/app/model"
)

type UserResource struct {
	FirstName string      `json:"first_name"`
	LastName  interface{} `json:"last_name"`
	Email     string      `json:"email"`
}

func NewUserResource(user *model.User) *UserResource {
	return &UserResource{
		FirstName: user.FirstName,
		Email:     user.Email,
		LastName:  user.LastName.String,
	}
}

func NewUsersResource(users *[]model.User) []UserResource {
	var usersResponse []UserResource
	for _, user := range *users {
		usersResponse = append(usersResponse, *NewUserResource(&user))
	}

	return usersResponse
}
