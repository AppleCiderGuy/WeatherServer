package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherRequestResponse struct {
	Requested_city string
	Message        string
}

func WeatherRequestHandler(writer http.ResponseWriter, requestPointer *http.Request) error {

	requestPointer.ParseForm()
	city := requestPointer.Form.Get("city")
	fmt.Println("City is : ", city)
	if city == "" {
		return fmt.Errorf("need query parameter 'city'")
	}

	response := WeatherRequestResponse{
		Requested_city: city,
		Message:        "Root Url under Construction",
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)

	return nil
}
