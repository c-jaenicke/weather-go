package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"weather-go/pkg/env"
	"weather-go/pkg/request"
)

// GetWeatherData returns api response containing weather data
func GetWeatherData(location string) (ApiResponse, string) {
	var apiKey = getApiKey()
	lat, lon, name := geocodeLocation(location)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=%f&lon=%f&exclude=minutely,daily,alerts&appid=%s&units=metric", lat, lon, apiKey)

	responseData := request.Request(url)
	var responseObject ApiResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject, name
}

// geocodeLocation uses api to turn location into coordinates and name
func geocodeLocation(location string) (lat float64, lon float64, name string) {
	var apiKey = getApiKey()

	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", location, apiKey)
	responseData := request.Request(url)

	var responseObject GeocodeResponse
	json.Unmarshal(responseData, &responseObject)

	if len(responseObject) == 0 {
		fmt.Println("No valid location found!")
		os.Exit(1)
	}

	return responseObject[0].Lat, responseObject[0].Lon, responseObject[0].Name
}

func getApiKey() string {
	apiKey, err := env.GetEnv("API_KEY")
	if err != nil {
		log.Fatal(err.Error())
	}

	return apiKey
}
