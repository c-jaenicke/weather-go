package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"weather-go/pkg/env"
	"weather-go/pkg/weather"
)

func main() {
	var location string
	var mode string
	var pathEnv string
	flag.StringVar(&location, "location", "New York", "Set location to get weather")
	flag.StringVar(&mode, "mode", "small", "Set output mode\nsmall: Print single line consisting of location name, temperature, weather\nfull: Prints full data of current weather")
	flag.StringVar(&pathEnv, "env", "./.env", "Path to .env file containing api key and optional location")
	flag.Parse()

	// give path to env module
	env.Path = pathEnv

	// get location from .env file if ".env" given as location
	if location == ".env" {
		var err error
		location, err = env.GetEnv("LOCATION")
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	weatherData, locationName := weather.GetWeatherData(location)
	var data = ""

	switch mode {
	case "small":
		currentWeather := fmt.Sprintf("%s: %s°C %s", locationName, strconv.FormatFloat(weatherData.Current.Temp, 'f', -1, 64), weatherData.Current.Weather[0].Description)
		data = currentWeather

	case "full":
		currentWeather := fmt.Sprintf("%s: \n\t%s, %s\n\tTemperature: %s °C\n\tHumidity: %d hpa\n\tPressure: %d\n\tWind: %f m/s from %d", locationName, weatherData.Current.Weather[0].Main, weatherData.Current.Weather[0].Description, strconv.FormatFloat(weatherData.Current.Temp, 'f', -1, 64), weatherData.Current.Humidity, weatherData.Current.Pressure, weatherData.Current.WindSpeed, weatherData.Current.WindDeg)
		data = currentWeather

	default:
		data = "Invalid Mode"
	}

	fmt.Println(data)
}
