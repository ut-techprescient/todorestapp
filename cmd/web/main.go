package main

import (
	"net/http"
	"os"

	"todorestapp.ut/cmd/config"
)

func main() {

	addr := ":4000"

	// configure logging into application
	infoLog, _ := config.CreateLogger(os.Stdout, "info")
	errorLog, _ := config.CreateLogger(os.Stdout, "error")

	// initializing app exceptions

	// initializing all environment variables
	appConfig := &config.AppEnv{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		AppError: &config.AppError{ErrorLog: errorLog},
	}

	// init server configuration
	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  routes(appConfig),
	}

	infoLog.Printf("Starting server at addr: %s", addr)

	err := srv.ListenAndServe()

	errorLog.Fatal(err)

}
