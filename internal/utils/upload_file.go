package utils

import (
	"io"
	"net/http"
	"os"
)

func uploadFile(r *http.Request) (string, error) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")

	if err != nil {
		return "", err
	}

	defer file.Close()

	fileName := handler.Filename

	dst, err := os.Create("tmp/" + fileName)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return dst.Name(), nil

}

func UploadFile(r *http.Request) (string, error) {
	return uploadFile(r)
}
