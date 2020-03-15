package app

import (
	"errors"
	"fileService/pkg/crud/services/files"
	"net/http"
)

type server struct {
	router        http.Handler
	filesSvc      *files.FilesSvc
	templatesPath string
	assetsPath    string
	mediaPath     string
}

func NewServer(router http.Handler, filesSvc *files.FilesSvc, templatesPath string, assetsPath string, mediaPath string) *server {
	if filesSvc == nil {
		panic(errors.New("filesSvc can't be nil"))
	}
	if templatesPath == "" {
		panic(errors.New("templatesPath can't be empty"))
	}
	if assetsPath == "" {
		panic(errors.New("assetsPath can't be empty"))
	}
	if mediaPath == "" {
		panic(errors.New("mediaPath can't be empty"))
	}

	return &server{
		router:        router,
		filesSvc:      filesSvc,
		templatesPath: templatesPath,
		assetsPath:    assetsPath,
		mediaPath:     mediaPath,
	}
}

func (receiver *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	receiver.router.ServeHTTP(writer, request)
}