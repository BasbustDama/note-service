package server

import "net/http"

type HttpServer struct {
	Router http.Handler
	Port   string
}

func New(router http.Handler, port string) *HttpServer {
	return &HttpServer{
		Router: router,
		Port:   port,
	}
}

func (server *HttpServer) Run() {
	err := http.ListenAndServe(server.Port, server.Router)
	if err != nil {
		panic(err.Error())
	}
}
