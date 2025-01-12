package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-restful/app/middleware"
	"github.com/go-restful/app/response"
	"github.com/go-restful/app/router"
	"github.com/go-restful/helper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(15)

	if err != nil {
		panic(err)
	}

	router := router.NewRouter(db)

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		response.JsonResponse(w, response.NewInternalServerError("Something went wrong!", err))
	}

	server := &http.Server{
		Addr:    ":5000",
		Handler: middleware.Logging(middleware.DBMiddleware(router)),
	}

	fmt.Println("Server is running on :5000")
	helper.ErrorPanic(server.ListenAndServe())
}
