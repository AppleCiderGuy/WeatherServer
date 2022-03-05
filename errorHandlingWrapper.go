package main

import (
	"net/http"
)

type customHandlerFunctions func(writer http.ResponseWriter, requestPointer *http.Request) error

func (fn customHandlerFunctions) ServeHTTP(writer http.ResponseWriter, requestPointer *http.Request) {
	err := fn(writer, requestPointer)
	if err == nil {
		return
	}

	//error handling
	writer.WriteHeader(400)
	writer.Write([]byte(err.Error()))

	return
}
