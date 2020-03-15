package files

import (
	"errors"
	"fmt"
	jsonwriter "github.com/shuhrat-shokirov/rest/pkg/rest"
	"github.com/google/uuid"
	"io"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

type FilesSvc struct {
	mediaPath string
}

func NewFilesSvc(mediaPath string) *FilesSvc {
	if mediaPath == "" {
		panic(errors.New("media path can't be nil")) // <- be accurate
	}
	return &FilesSvc{mediaPath: mediaPath}
}

func (receiver *FilesSvc) Save(sources io.Reader, contentType string) (name string, err error) {
	var path string
	extensions, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", err
	}
	if len(extensions) == 0 {
		return "", errors.New("invalid extension")
	}
	uuidV4 := uuid.New().String()
	name = fmt.Sprintf("%s%s", uuidV4, extensions[0])
	path = filepath.Join(receiver.mediaPath, name)
	dst, err := os.Create(path)
	if err != nil {
		log.Print("can't close file")
	}
	defer func() {
		if dst.Close() != nil {
			log.Print("can't close file")
		}
	}()
	_, err = io.Copy(dst, sources)
	if err != nil {
		log.Printf("ca't save file: %v", sources)
	}
	filesPath := strings.Split(path, name)
	pathFile := filesPath[0]
	upload, err := jsonwriter.JsonFileUpload(pathFile)
	return upload, nil
}