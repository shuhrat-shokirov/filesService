package main

// package
// import
// var + type
// method + function

import (
	"fileService/cmd/crud/app"
	"fileService/pkg/crud/services/files"
	"flag"
	"net"
	"os"
	"path/filepath"
)

var (
	host = flag.String("host", "", "Server host")
	port = flag.String("port", "", "Server port")
)

const envHost = "HOST"
const envPort = "PORT"

func fromFLagOrEnv(flag *string, envName string) (server string, ok bool){
	if *flag != ""{
		return *flag, true
	}
	return os.LookupEnv(envName)
}

func main() {
	flag.Parse()
	hostf, ok := fromFLagOrEnv(host, envHost)
	if !ok {
		hostf = *host
	}
	portf, ok := fromFLagOrEnv(port, envPort)
	if !ok {
		portf = *port
	}
	addr := net.JoinHostPort(hostf, portf)
	start(addr)
}

func start(addr string) {
	router := app.NewExactMux()
	templatesPath := filepath.Join("web", "templates")
	assetsPath := filepath.Join("web", "assets")
	mediaPath := filepath.Join("web", "media")
	filesSvc := files.NewFilesSvc(mediaPath)
	server := app.NewServer(
		router,
		filesSvc,
		templatesPath,
		assetsPath,
		mediaPath,
	)
	server.GorillaInit(addr)
	//panic(http.ListenAndServe(addr, server))
}