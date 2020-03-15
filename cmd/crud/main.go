package main

// package
// import
// var + type
// method + function

import (
	"context"
	"fileService/cmd/crud/app"
	"fileService/pkg/crud/services/files"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"net"
	"os"
	"path/filepath"
)

var (
	host = flag.String("host", "", "Server host")
	port = flag.String("port", "", "Server port")
	dsn  = flag.String("dsn", "", "Postgres DSN")
)

const envHost = "HOST"
const envPort = "PORT"
const envDSN  = "DATABASE_URL"

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
	dsnf, ok := fromFLagOrEnv(dsn, envDSN)
	if !ok {
		dsnf = *dsn
	}

	addr := net.JoinHostPort(hostf, portf)
	start(addr, dsnf)
}

func start(addr string, dsn string) {
	router := app.NewExactMux()
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	templatesPath := filepath.Join("web", "templates")
	assetsPath := filepath.Join("web", "assets")
	mediaPath := filepath.Join("web", "media")
	filesSvc := files.NewFilesSvc(mediaPath)
	server := app.NewServer(
		router,
		pool,
		filesSvc,
		templatesPath,
		assetsPath,
		mediaPath,
	)
	server.GorillaInit(addr)
	//panic(http.ListenAndServe(addr, server))
}