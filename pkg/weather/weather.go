package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"weather-go/pkg/env"
	"weather-go/pkg/request"
)

// GetWeatherData returns api response containing weather data
func GetWeatherData(location string) (*CurrentWeather, string, error) {
	apiKey := getApiKey()
	lat, lon, name, err := geocodeLocation(location)
	if err != nil {
		log.Println("GetWeatherData: failed to geocode: " + err.Error())
		return nil, "", err
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&exclude=minutely,daily,alerts&appid=%s&units=metric", lat, lon, apiKey)
	responseData := request.Request(url)

	var responseObject CurrentWeather
	jsonError := json.Unmarshal(responseData, &responseObject)
	if jsonError != nil {
		log.Fatalf("GetWeatherData: failed to json unmarshal response: " + jsonError.Error())
		return nil, "", err
	}

	return &responseObject, name, nil
}

// GetForecast retrieves the weather forecast for a given location.
func GetForecast(location string) (*ForecastResponse, string, error) {
	apiKey := getApiKey()
	lat, lon, name, err := geocodeLocation(location)
	if err != nil {
		log.Println("GetForecast: failed to geocode: " + err.Error())
		return nil, "", err
	}

	// https://openweathermap.org/forecast5
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=metric&cnt=10", lat, lon, apiKey)
	responseData := request.Request(url)

	var responseObject ForecastResponse
	jsonError := json.Unmarshal(responseData, &responseObject)
	if jsonError != nil {
		log.Fatalf("GetForecast: failed to json unmarshal response: " + jsonError.Error())
		return nil, "", jsonError
	}

	return &responseObject, name, nil
}

// geocodeLocation uses api to turn location into coordinates
func geocodeLocation(location string) (lat float64, lon float64, name string, err error) {
	apiKey := getApiKey()
	location = strings.TrimSpace(location)

	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", location, apiKey)
	responseData := request.Request(url)

	var responseObject GeocodeResponse
	json.Unmarshal(responseData, &responseObject)

	if len(responseObject) == 0 {
		return 0, 0, "", fmt.Errorf("failed to find valid location")
	} else {
		return responseObject[0].Lat, responseObject[0].Lon, responseObject[0].Name + ", " + responseObject[0].Country, nil
	}
}

// getApiKey get api key from .env file
func getApiKey() string {
	apiKey, err := env.GetEnv("API_KEY")
	if err != nil {
		log.Fatal(err.Error())
	}

	return apiKey
}
