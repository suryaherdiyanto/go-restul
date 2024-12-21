package repository

import (
	"context"

	"github.com/go-restful/app/model"
	"github.com/go-restful/app/request"
)

type PostRepository interface {
	FindById(ctx context.Context, id int) (model.Post, bool)
	FindBy(ctx context.Context, field string, value interface{}) (model.Post, bool)
	All(ctx context.Context) []model.Post
	Create(ctx context.Context, data *request.PostRequest) model.Post
	Update(ctx context.Context, id int, data *request.PostUpdateRequest) model.Post
	Delete(ctx context.Context, id int)
}
