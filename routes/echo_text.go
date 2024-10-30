package routes

import "net/http"

type EchoTextHandler struct {
}

func NewEchoTextHandler() *EchoTextHandler {
	return &EchoTextHandler{}
}

func (e *EchoTextHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, Text!"))
}

func (e *EchoTextHandler) Pattern() string {
	return "/echo-text"
}
