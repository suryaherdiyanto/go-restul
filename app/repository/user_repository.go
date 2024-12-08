package repository

import (
	"github.com/go-restful/app/model"
	"github.com/go-restful/app/request"
)

type UserRepository interface {
	FindById(id int) (model.User, bool)
	All() []model.User
	Create(data *request.UserRequest) model.User
	Update(id int, data *request.UserRequest) (model.User, bool)
	Delete(id int)
}
