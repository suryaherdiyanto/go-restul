package repository

import (
	"context"

	"github.com/go-restful/app/model"
	"github.com/go-restful/app/request"
)

type CategoryRepository interface {
	FindById(ctx context.Context, id int) (model.Category, bool)
	FindBy(ctx context.Context, field string, value interface{}) (model.Category, bool)
	All(ctx context.Context) []model.Category
	Create(ctx context.Context, data *request.CategoryRequest) model.Category
	Update(ctx context.Context, id int, data *request.CategoryUpdateRequest) model.Category
	Delete(ctx context.Context, id int)
}
