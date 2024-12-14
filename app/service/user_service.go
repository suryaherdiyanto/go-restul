package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-restful/app/model"
	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/request"
	"github.com/go-restful/helper"
)

type UserService struct {
	DB *sql.DB
}

func (userService *UserService) All(ctx context.Context) []model.User {
	q := "select * from users"

	row, err := userService.DB.QueryContext(ctx, q)
	helper.ErrorPanic(err)

	var users []model.User
	for row.Next() {
		var user model.User
		row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt)

		users = append(users, user)
	}

	return users
}

func (userService *UserService) FindById(ctx context.Context, id int) (model.User, bool) {
	q := "select * from users where id = ? limit 1"

	row, err := userService.DB.QueryContext(ctx, q, id)
	helper.ErrorPanic(err)

	if row.Next() {
		var user model.User
		row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt)

		return user, true
	}

	return model.User{}, false
}

func (userService *UserService) Create(ctx context.Context, data *request.UserRequest) model.User {
	q := "insert into users(first_name,last_name,email,created_at,updated_at) values(?, ?, ?, now(), now())"

	row, err := userService.DB.ExecContext(ctx, q, data.FirstName, data.LastName, data.Email)
	helper.ErrorPanic(err)

	id, err := row.LastInsertId()
	helper.ErrorPanic(err)

	user := model.User{}
	user.Id = int(id)
	user.Email = data.Email
	user.FirstName = data.FirstName
	user.LastName = helper.HandleNullString(data.LastName)
	user.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	user.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return user
}

func (userService *UserService) Update(ctx context.Context, id int, data *request.UserRequest) model.User {

	q := "update users set first_name=?, last_name=?, email=?, updated_at=now() where id = ?"

	_, err := userService.DB.ExecContext(ctx, q, data.FirstName, data.LastName, data.Email, id)
	helper.ErrorPanic(err)

	user := model.User{}

	user.Id = id
	user.Email = data.Email
	user.FirstName = data.FirstName
	user.LastName = helper.HandleNullString(data.LastName)
	user.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	user.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return user
}

func (userService *UserService) Delete(ctx context.Context, id int) {
	q := "delete from users where id = ?"

	_, err := userService.DB.ExecContext(ctx, q, id)
	helper.ErrorPanic(err)

}

func NewUserService(db *sql.DB) repository.UserRepository {
	return &UserService{
		DB: db,
	}
}
