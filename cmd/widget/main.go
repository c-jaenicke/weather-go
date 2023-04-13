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
	var emojis bool
	flag.StringVar(&location, "location", "", "Set location to get weather.\nExpected format: <city_name,state_code,country_code>\nCan use '.env' as value, doing so will load the location from the given .env file")
	flag.StringVar(&mode, "mode", "small", "Set output mode\nsmall: Print single line consisting of location name, temperature, weather\nfull: Prints full data of current weather\nforecast: Prints a 24 hour forecast for the location")
	flag.StringVar(&pathEnv, "env", "", "Path to .env file containing api key and optional location")
	flag.BoolVar(&emojis, "emojis", false, "Enable icons, disabled by default. Requires an emoji font to be installed and enabled!")
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
		var currentWeather string
		if emojis {
			currentWeather = fmt.Sprintf("%s: %s °C %s %s", locationName, shortenFloat(weatherData.Current.Temp), weatherData.Current.Weather[0].Description, weather.IconMap[weatherData.Current.Weather[0].ID])
		} else {
			currentWeather = fmt.Sprintf("%s: %s °C %s", locationName, shortenFloat(weatherData.Current.Temp), weatherData.Current.Weather[0].Description)
		}

		data = currentWeather

	case "full":
		currentWeather := fmt.Sprintf("%s:"+
			"\n%s, %s"+
			"\n\tTemperature: %s °C"+
			"\n\tHumidity: %d%%"+
			"\n\tPressure: %d hpa"+
			"\n\tWind: %s m/s from %s",
			locationName,
			weatherData.Current.Weather[0].Main, weatherData.Current.Weather[0].Description,
			shortenFloat(weatherData.Current.Temp),
			weatherData.Current.Humidity,
			weatherData.Current.Pressure,
			shortenFloat(weatherData.Current.WindSpeed), weather.WindDegreesToDirection(weatherData.Current.WindDeg))

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
				"\n\tWind: %s m/s from %s"+
				"\n\n",
				unixToDateTime(item.Dt),
				item.Weather[0].Description,
				shortenFloat(item.Main.Temp),
				item.Main.Humidity,
				item.Main.Pressure,
				shortenFloat(item.Wind.Speed), weather.WindDegreesToDirection(item.Wind.Deg))
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
