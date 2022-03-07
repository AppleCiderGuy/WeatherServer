package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PointWeatherRequestResponse struct {
	Message string
}

func PointWeatherRequestHandler(responseWriter http.ResponseWriter, requestPointer *http.Request) error {

	fmt.Println("starting request to get weather at coordinates")

	requestPointer.ParseForm()
	lat := requestPointer.Form.Get("lat")
	long := requestPointer.Form.Get("long")

	if lat == "" || long == "" {
		return fmt.Errorf("request needs query params 'lat' and 'long'")
	}

	requestUrl := fmt.Sprintf("https://api.weather.gov/points/%s,%s", lat, long)
	fmt.Println("requesting weather from url: ", requestUrl)
	resp, err := http.Get(requestUrl)

	if err != nil {
		fmt.Println("request error: ", err.Error())
		return fmt.Errorf("Internal Failure: unable to fetch weather for providede coordinates(%v,%v)", lat, long)
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("parsing error: ", err.Error())
		return fmt.Errorf("Internal Failure: unable to fetch weather for providede coordinates(%v,%v)", lat, long)
	}

	json.NewEncoder(responseWriter).Encode(PointWeatherRequestResponse{Message: fmt.Sprint("weather in RAW form :", string(respBody))})

	return nil
}
