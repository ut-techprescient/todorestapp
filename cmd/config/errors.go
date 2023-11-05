package config

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type AppError struct {
	ErrorLog *log.Logger
}

func (appError *AppError) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	appError.ErrorLog.Print(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (appError *AppError) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (appError *AppError) NotFound(w http.ResponseWriter) {
	appError.ClientError(w, http.StatusNotFound)
}

func (appError *AppError) MethodNotAllowed(w http.ResponseWriter) {
	appError.ClientError(w, http.StatusMethodNotAllowed)
}
