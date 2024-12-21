package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-restful/app/model"
	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/request"
	"github.com/go-restful/helper"
)

type PostService struct {
	DB        *sql.DB
	TableName string
}

func (postService *PostService) All(ctx context.Context) []model.Post {
	q := fmt.Sprintf("select * from %s", postService.TableName)

	row, err := postService.DB.QueryContext(ctx, q)
	helper.ErrorPanic(err)

	var posts []model.Post

	for row.Next() {
		var post model.Post

		row.Scan(&post.Id, &post.Title, &post.Category, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt)

		posts = append(posts, post)
	}

	return posts
}

func (postService *PostService) FindById(ctx context.Context, id int) (model.Post, bool) {
	q := fmt.Sprintf("select * from %s where id = ? limit 1", postService.TableName)

	row, err := postService.DB.QueryContext(ctx, q, id)
	helper.ErrorPanic(err)

	if row.Next() {
		var post model.Post
		row.Scan(&post.Id, &post.Title, &post.Category, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt)

		return post, true
	}

	return model.Post{}, false
}

func (postService *PostService) FindBy(ctx context.Context, field string, value interface{}) (model.Post, bool) {
	q := fmt.Sprintf("select * from %s where %s = ? limit 1", postService.TableName, field)

	row, err := postService.DB.QueryContext(ctx, q, value)
	helper.ErrorPanic(err)

	if row.Next() {
		var post model.Post
		row.Scan(&post.Id, &post.Title, &post.Category, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt)

		return post, true
	}

	return model.Post{}, false
}

func (postService *PostService) Create(ctx context.Context, data *request.PostRequest) model.Post {
	q := fmt.Sprintf("insert into %s (title, category, user_id, content, created_at, updated_at) values (?, ?, ?, ?, now(), now())", postService.TableName)

	row, err := postService.DB.ExecContext(ctx, q, data.Title, data.Category, data.UserID, data.Content)
	helper.ErrorPanic(err)

	id, err := row.LastInsertId()
	helper.ErrorPanic(err)

	post := model.Post{}
	post.Id = int(id)
	post.Title = data.Title
	post.Category = data.Category
	post.UserID = data.UserID
	post.Content = data.Content
	post.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	post.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return post
}

func (postService *PostService) Update(ctx context.Context, id int, data *request.PostUpdateRequest) model.Post {

	q := fmt.Sprintf("update %s set title = ?, category = ?, user_id = ?, content = ?, updated_at = now() where id = ?", postService.TableName)

	_, err := postService.DB.ExecContext(ctx, q, data.Title, data.Category, data.UserID, data.Content, id)
	helper.ErrorPanic(err)

	post := model.Post{}
	post.Id = int(id)
	post.Title = data.Title
	post.Category = data.Category
	post.UserID = data.UserID
	post.Content = data.Content
	post.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	post.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return post
}

func (postService *PostService) Delete(ctx context.Context, id int) {
	q := fmt.Sprintf("delete from %s where id = ?", postService.TableName)

	_, err := postService.DB.ExecContext(ctx, q, id)
	helper.ErrorPanic(err)

}

func NewPostService(db *sql.DB) repository.PostRepository {
	return &PostService{DB: db, TableName: "posts"}
}
