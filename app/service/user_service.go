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

func (userService *UserService) All() []model.User {
	q := "select * from users"
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

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

func (userService *UserService) FindById(id int) (model.User, bool) {
	q := "select * from users where id = ? limit 1"
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	row, err := userService.DB.QueryContext(ctx, q, id)
	helper.ErrorPanic(err)

	if row.Next() {
		var user model.User
		row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt)

		return user, true
	}

	return model.User{}, false
}

func (userService *UserService) Create(data *request.UserRequest) model.User {
	q := "insert into users(first_name,last_name,email,created_at,updated_at) values(?, ?, ?, now(), now())"
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	row, err := userService.DB.ExecContext(ctx, q, data.FirstName, data.LastName, data.Email)
	helper.ErrorPanic(err)

	id, err := row.LastInsertId()
	helper.ErrorPanic(err)

	user, _ := userService.FindById(int(id))

	return user
}

func (userService *UserService) Update(id int, data *request.UserRequest) (model.User, bool) {
	user, ok := userService.FindById(id)

	if !ok {
		return model.User{}, false
	}

	q := "update users set first_name=?, last_name=?, email=?, updated_at=now() where id = ?"
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	_, err := userService.DB.ExecContext(ctx, q, data.FirstName, data.LastName, data.Email, user.Id)
	helper.ErrorPanic(err)

	user.Email = data.Email
	user.FirstName = data.FirstName
	user.LastName = data.LastName.(sql.NullString)

	return user, true
}

func (userService *UserService) Delete(id int) {
	q := "delete from users where id = ?"
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	_, err := userService.DB.ExecContext(ctx, q, id)
	helper.ErrorPanic(err)

}

func NewUserService(db *sql.DB) repository.UserRepository {
	return &UserService{
		DB: db,
	}
}
