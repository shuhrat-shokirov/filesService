package app

const multipartMaxBytes = 10 * 1024 * 1024

func (receiver *server) Init(addr string) {

	receiver.router.POST("/api/files", receiver.handleFilesSave())
	//router.HandleFunc("POST", uploading, receiver.handleFilesSave())

	receiver.router.GET(upload, receiver.handleGetFile())

	//http.Handle(grut, router)
}
