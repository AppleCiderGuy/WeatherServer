package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PointWeatherRequestResponse struct {
	Message  string
	Forecast Forecast
}

func PointWeatherRequestHandler(responseWriter http.ResponseWriter, requestPointer *http.Request) error {

	fmt.Println("starting request to get weather at coordinates")

	requestPointer.ParseForm()
	lat := requestPointer.Form.Get("lat")
	long := requestPointer.Form.Get("long")

	if lat == "" || long == "" {
		return fmt.Errorf("request needs query params 'lat' and 'long'")
	}

	forecast, err := getWeather(lat, long)

	if err != nil {
		return fmt.Errorf("Internal Failure: unable to fetch weather for providede coordinates(%v,%v) \n -> %s", lat, long, err.Error())
	}

	json.NewEncoder(responseWriter).Encode(PointWeatherRequestResponse{Message: "-", Forecast: forecast})

	return nil
}

func getWeather(lat string, long string) (Forecast, error) {

	requestUrl := fmt.Sprintf("https://api.weather.gov/points/%s,%s", lat, long)
	fmt.Println("requesting weather from url: ", requestUrl)
	resp, err := http.Get(requestUrl)

	if err != nil {
		fmt.Println("request error: ", err.Error())
		return Forecast{}, fmt.Errorf("Internal Failure: unable to fetch weather for providede coordinates(%v,%v)", lat, long)
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	var weatherAPIResponse WeatherAPIResponse
	err1 := json.Unmarshal([]byte(respBody), &weatherAPIResponse)

	if err1 != nil {
		fmt.Println("Error Parsing response from weather API", err1.Error())
		return Forecast{}, fmt.Errorf("Unable to parse response")
	}

	forecast, err := getForecast(weatherAPIResponse.Properties.Forecast)
	if err != nil {
		return Forecast{}, err
	}

	return forecast, nil
}

func getForecast(url string) (Forecast, error) {
	fmt.Printf("Fetching Forecast from url %s", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching Forecast from url %s : %s", url, err.Error())
		return Forecast{}, fmt.Errorf("Error fetching Forast")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	var forecast ForecastProperties
	err1 := json.Unmarshal([]byte(respBody), &forecast)

	if err1 != nil {
		fmt.Printf("Error Parsing Forecast fetched from url: %s , Error: %s", url, err1.Error())
		return Forecast{}, fmt.Errorf("Error parsing Forast")
	}

	return forecast.Properties, nil
}

type WeatherAPIResponse struct {
	Properties properties `json:"properties"`
}

type properties struct {
	Forecast string `json:"forecast"`
}

type ForecastProperties struct {
	Properties Forecast `json:"properties"`
}

type Forecast struct {
	Periods []PerodicTempratures `json:"periods"`
}

type PerodicTempratures struct {
	Name     string `json:"name"`
	Temp     int    `json:"temperature"`
	TempUnit string `json:"temperatureUnit"`
	Short    string `json:"shortForecast"`
	Details  string `json:"detailedForecast"`
}
