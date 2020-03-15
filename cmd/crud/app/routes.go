package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const multipartMaxBytes = 10 * 1024 * 1024

func (receiver *server) GorillaInit() {
	router := mux.NewRouter()
	router.HandleFunc(grut, receiver.handleRedirect())
	router.HandleFunc(upload, receiver.handleUpload())
	router.HandleFunc(uploading, receiver.handleFilesSave())
	router.HandleFunc(favicon, receiver.handleFavicon())
	router.HandleFunc(media, http.StripPrefix(media, http.FileServer(http.Dir(receiver.mediaPath))).ServeHTTP)
	router.PathPrefix(media + grut).Handler(http.StripPrefix(media + grut, http.FileServer(http.Dir(receiver.mediaPath))))
	router.PathPrefix(grut).HandlerFunc(receiver.handleGetFile())
	http.Handle(grut, router)
	fmt.Println("Server is listening...")
	err := http.ListenAndServe(PortGor, nil)
	if err != nil {
		log.Fatal("can't start server")
	}
}
