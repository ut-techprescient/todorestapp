package config

import (
	"errors"
	"io"
	"log"
)

var logPrefix = map[string]string{
	"info":  "INFO::\t",
	"error": "ERROR::\t",
}

func CreateLogger(out io.Writer, prefix string) (*log.Logger, error) {
	logLevel, ok := logPrefix[prefix]
	if !ok {
		return nil, errors.New("WRONG LEVEL valid levels [info, error]")
	}
	if prefix == "info" {
		return log.New(out, logLevel, log.Ldate|log.Ltime|log.Lshortfile), nil
	}
	return log.New(out, logLevel, log.Ldate|log.Ltime|log.Llongfile), nil
}
