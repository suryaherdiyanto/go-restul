package main

import (
	"database/sql"
	"net/http"

	"github.com/go-restful/app/controller"
	"github.com/go-restful/app/service"
	"github.com/go-restful/helper"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/gorestful")

	helper.ErrorPanic(err)

	router := httprouter.New()

	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)

	router.GET("/api/users", userController.Index)

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	helper.ErrorPanic(server.ListenAndServe())
}
