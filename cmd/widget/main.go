package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
	"weather-go/pkg/env"
	"weather-go/pkg/weather"
)

func main() {
	var location string
	var mode string
	var pathEnv string
	flag.StringVar(&location, "location", "", "Set location to get weather")
	flag.StringVar(&mode, "mode", "small", "Set output mode\nsmall: Print single line consisting of location name, temperature, weather\nfull: Prints full data of current weather")
	flag.StringVar(&pathEnv, "env", "", "Path to .env file containing api key and optional location")
	flag.Parse()

	if pathEnv == "" {
		log.Fatal("Missing path to .env file! Exiting ...")
	}

	if location == "" {
		log.Fatal("No location given! Exiting ...")
	}

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

	weatherData, locationName, err := weather.GetWeatherData(location)
	if err != nil {
		log.Fatal(err.Error())
	}
	var data = ""

	switch mode {
	case "small":
		currentWeather := fmt.Sprintf("%s: %s °C %s", locationName, shortenFloat(weatherData.Current.Temp), weatherData.Current.Weather[0].Description)
		data = currentWeather

	case "full":
		currentWeather := fmt.Sprintf("%s:"+
			"\n%s, %s"+
			"\n\tTemperature: %s °C"+
			"\n\tHumidity: %d%%"+
			"\n\tPressure: %d hpa"+
			"\n\tWind: %f m/s from %d", locationName, weatherData.Current.Weather[0].Main, weatherData.Current.Weather[0].Description, shortenFloat(weatherData.Current.Temp), weatherData.Current.Humidity, weatherData.Current.Pressure, weatherData.Current.WindSpeed, weatherData.Current.WindDeg)

		data = currentWeather

	case "forecast":
		forecast, name, err := weather.GetForecast(location)
		if err != nil {
			log.Println(err.Error())
		}

		message := fmt.Sprintf("Forecast for %s:\n", name)

		for _, item := range forecast.List {
			message += fmt.Sprintf("\t%s"+
				"\n%s"+
				"\n\tTemp: %s °C"+
				"\n\tHumidity: %d%%"+
				"\n\tPressure: %d hpa"+
				"\n\tWind: %s m/s from %d"+
				"\n\n", unixToDateTime(item.Dt), item.Weather[0].Description, shortenFloat(item.Main.Temp), item.Main.Humidity, item.Main.Pressure, shortenFloat(item.Wind.Speed), item.Wind.Deg)
		}

		data = message

	default:
		data = "Invalid Mode"
	}

	fmt.Println(data)
}

// unixToDateTime transforms unix time as int to a datetime string, formatted as "YYYY-MM-DD HH-MM UTC-OFFSET TIMEZONE"
func unixToDateTime(unix int) string {
	tm := time.Unix(int64(unix), 0)
	return tm.String()
}

// shortenFloat transforms the given float64 value into a string with 2 decimal values
func shortenFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
