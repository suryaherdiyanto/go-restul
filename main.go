package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-restful/app/controller"
	"github.com/go-restful/app/service"
	"github.com/go-restful/helper"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/gorestful")
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(15)

	helper.ErrorPanic(err)

	router := httprouter.New()

	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)

	router.GET("/api/users", userController.Index)
	router.GET("/api/users/:id", userController.Show)
	router.POST("/api/users", userController.Store)
	router.PUT("/api/users/:id/update", userController.Update)
	router.DELETE("/api/users/:id/delete", userController.Delete)

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	fmt.Println("Server is running on :5000")
	helper.ErrorPanic(server.ListenAndServe())
}
