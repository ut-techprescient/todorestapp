package config

import "log"

type AppEnv struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	AppError *AppError
}
