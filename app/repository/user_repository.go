package repository

import (
	"context"

	"github.com/go-restful/app/model"
	"github.com/go-restful/app/request"
)

type UserRepository interface {
	FindById(ctx context.Context, id int) (model.User, bool)
	FindBy(ctx context.Context, field string, value interface{}) (model.User, bool)
	All(ctx context.Context) []model.User
	Create(ctx context.Context, data *request.UserRequest) model.User
	Update(ctx context.Context, id int, data *request.UserUpdateRequest) model.User
	Delete(ctx context.Context, id int)
}
