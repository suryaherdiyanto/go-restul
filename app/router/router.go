package router

import (
	"database/sql"

	"github.com/go-restful/app/controller"
	"github.com/go-restful/app/service"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(db *sql.DB) *httprouter.Router {
	router := httprouter.New()

	userService := service.NewUserService(db)
	postService := service.NewPostService(db)

	userController := controller.NewUserController(userService)
	postController := controller.NewPostController(postService)

	router.GET("/api/users", userController.Index)
	router.GET("/api/users/:id", userController.Show)
	router.POST("/api/users", userController.Store)
	router.PUT("/api/users/:id/update", userController.Update)
	router.DELETE("/api/users/:id/delete", userController.Delete)

	router.GET("/api/posts", postController.Index)
	router.GET("/api/posts/:id", postController.Show)
	router.POST("/api/posts", postController.Store)
	router.PUT("/api/posts/:id/update", postController.Update)
	router.DELETE("/api/posts/:id/delete", postController.Delete)

	return router
}
