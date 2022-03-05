package main

import (
	"fmt"
	"log"
	"net/http"
)

func setupEndpointRouting() {
	http.Handle("/report", customHandlerFunctions(WeatherRequestHandler))
}

func startServer() {
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Starting Server....")
	setupEndpointRouting()
	startServer()
}
