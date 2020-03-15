package app

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"html/template"
)
func(receiver *server) handleGetFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		fileName := strings.TrimPrefix(request.RequestURI, grut)
		log.Print(fileName)

		file, err := ioutil.ReadFile(filepath.Join(receiver.mediaPath, fileName))
		if err != nil {
			panic(err)
		}
		_, err = w.Write(file)
		if err != nil {
			log.Print(err)
		}
	}
}

func (receiver *server) handleRedirect() func(responseWriter http.ResponseWriter, request *http.Request) {

	return func(responseWriter http.ResponseWriter, request *http.Request) {

		http.Redirect(responseWriter, request, upload, http.StatusFound)
	}
}

func( receiver *server) handleUpload() func(http.ResponseWriter, *http.Request) {
	var (
		tpl *template.Template
		err error
	)

	tpl, err = template.ParseFiles(
		filepath.Join(receiver.templatesPath, "admin", fileHtml),
		filepath.Join(receiver.templatesPath, baseHtml),
	)
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {

		data := struct {
			Title string
		}{
			Title: "Uploader File",
		}

		err = tpl.Execute(writer, data)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	}
}

func( receiver *server) handleFilesSave() func(responseWriter http.ResponseWriter, request *http.Request) {

	return func(responseWriter http.ResponseWriter, request *http.Request) {

		err := request.ParseMultipartForm(multipartMaxBytes)
		if err != nil {
			log.Print(err)
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		_, header, err := request.FormFile(filesForm)
		uploadedFiles := ""
		if err != nil {
			log.Print(err)
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		contentType := header.Header.Get(conType)
		formFiles := request.MultipartForm
		files := formFiles.File

		for _, file := range files[filesForm] {
			openFile, err := file.Open()
			if err != nil {
				log.Printf("can't create file: %v", err)
			}

			uploadedFiles, err = receiver.filesSvc.Save(openFile, contentType)
			if err != nil {
				log.Printf("can't save file: %v", err)
			}

		}

		responseWriter.Header().Set(conType, value)
		_, err = responseWriter.Write([]byte(uploadedFiles))
		if err != nil { // ?
			http.Error(
				responseWriter,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
			return
		}

		http.Redirect(responseWriter, request, upload, http.StatusFound)
	}
}

func( receiver *server) handleFavicon() func(http.ResponseWriter, *http.Request) {
	file, err := ioutil.ReadFile(filepath.Join(receiver.assetsPath, favicon))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write(file)
		if err != nil {
			log.Print(err)
		}
	}
}