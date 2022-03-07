package main

import (
	"fmt"
	"log"
	"net/http"
)

func setupEndpointRouting() {
	http.Handle("/location", customHandlerFunctions(WeatherRequestHandler))
	http.Handle("/point", customHandlerFunctions(PointWeatherRequestHandler))
}

func startServer() {
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Starting Server....")
	setupEndpointRouting()
	startServer()
}
