package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"weather-go/internal/environment"
	"weather-go/internal/request"
)

// GetWeatherData returns API response containing weather data
func GetWeatherData(location string) (*CurrentWeather, string, error) {
	apiKey := getApiKey()
	lat, lon, name, err := geocodeLocation(location)
	if err != nil {
		return nil, "", fmt.Errorf("failed to geocode: %w", err)
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&exclude=minutely,daily,alerts&appid=%s&units=metric", lat, lon, apiKey)
	responseData, err := request.Request(url)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get weather data from API: %w", err)
	}

	var responseObject CurrentWeather
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		return nil, "", fmt.Errorf("failed to unmarshal weather data: %w", err)
	}

	return &responseObject, name, nil
}

// GetForecast retrieves the weather forecast for a given location.
func GetForecast(location string) (*ForecastResponse, string, error) {
	apiKey := getApiKey()
	lat, lon, name, err := geocodeLocation(location)
	if err != nil {
		return nil, "", fmt.Errorf("failed to geocode: %w", err)
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=metric&cnt=10", lat, lon, apiKey)
	responseData, err := request.Request(url)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get forecast data from API: %w", err)
	}

	var responseObject ForecastResponse
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		return nil, "", fmt.Errorf("failed to unmarshal forecast data: %w", err)
	}

	return &responseObject, name, nil
}

// geocodeLocation uses API to turn location into coordinates
func geocodeLocation(location string) (lat float64, lon float64, name string, err error) {
	apiKey := getApiKey()
	location = strings.ReplaceAll(location, " ", "")

	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", location, apiKey)
	responseData, err := request.Request(url)
	if err != nil {
		return 0, 0, "", fmt.Errorf("failed to get geocode location: %w", err)
	}

	var responseObject GeocodeResponse
	if err := json.Unmarshal(responseData, &responseObject); err != nil || len(responseObject) == 0 {
		return 0, 0, "", fmt.Errorf("failed to find valid location for '%s'", location)
	}

	return responseObject[0].Lat, responseObject[0].Lon, responseObject[0].Name + ", " + responseObject[0].Country, nil
}

// getApiKey gets API key from .env file
func getApiKey() string {
	apiKey, err := environment.GetEnv("API_KEY")
	if err != nil {
		log.Fatal("API key not found in environment: ", err)
	}
	return apiKey
}
