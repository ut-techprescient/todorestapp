package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"todorestapp.ut/cmd/config"
	"todorestapp.ut/internal/health"
	"todorestapp.ut/internal/middleware"
	"todorestapp.ut/internal/todo"
)

func routes(appConfig *config.AppEnv) http.Handler {
	mux := httprouter.New()
	// mux := &http.ServeMux{}

	mux.GET("/", health.HealthHandler(appConfig))

	mux.GET("/todo/list", todo.TodoTaskListHandler(appConfig))
	mux.POST("/todo/create", todo.TodoTaskCreateHandler(appConfig))
	mux.POST("/todo/upload", todo.TodoTaskUploadHandler(appConfig))
	mux.GET("/todo/get/:id", todo.TodoTaskGetHandler(appConfig))

	chainMiddleware := alice.New(middleware.RequestLoggerMiddleWare(appConfig))
	return chainMiddleware.Then(mux)
}
