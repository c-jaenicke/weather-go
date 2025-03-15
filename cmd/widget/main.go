package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"weather-go/internal/environment"
	"weather-go/internal/weather"
)

func main() {
	var location string
	var mode string
	var pathEnv string
	var emojis bool

	flag.StringVar(&location, "location", "", "Set location to get weather.\nExpected format: <city_name,state_code,country_code>")
	flag.StringVar(&mode, "mode", "small", "Set output mode\nsmall: Print single line consisting of location name, temperature, weather\nfull: Prints full data of current weather\nforecast: Prints a 24 hour forecast for the location")
	flag.StringVar(&pathEnv, "env", "", "Path to .env file containing api key and optional location (optional)")
	flag.BoolVar(&emojis, "emojis", false, "Enable icons, disabled by default. Requires an emoji font to be installed and enabled!")
	flag.Parse()

	if pathEnv == "" {
		pathEnv = ".env"
	}

	if err := environment.LoadEnv(pathEnv); err != nil {
		log.Fatalf("error loading .env file from path %s: %v", pathEnv, err)
	}

	if location == "" {
		var err error
		location, err = environment.GetEnv("LOCATION")
		if err != nil || location == "" {
			log.Fatal("no location specified and no LOCATION found in the .env file! Exiting...")
		}
	}

	weatherData, locationName, err := weather.GetWeatherData(location)
	if err != nil {
		log.Fatal("error getting weather data")
	}

	var data string
	switch mode {
	case "small":
		var currentWeather string
		if emojis {
			currentWeather = fmt.Sprintf("%s: %s °C %s %s", locationName, shortenFloat(weatherData.Main.Temp), weatherData.Weather[0].Description, weather.IconMap[weatherData.Weather[0].ID])
		} else {
			currentWeather = fmt.Sprintf("%s: %s °C %s", locationName, shortenFloat(weatherData.Main.Temp), weatherData.Weather[0].Description)
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
			weatherData.Weather[0].Main,
			weatherData.Weather[0].Description,
			shortenFloat(weatherData.Main.Temp),
			weatherData.Main.Humidity,
			weatherData.Main.Pressure,
			shortenFloat(weatherData.Wind.Speed),
			weather.WindDegreesToDirection(weatherData.Wind.Deg))

		data = currentWeather

	case "forecast":
		forecast, name, err := weather.GetForecast(location)
		if err != nil {
			log.Println("error getting forecast")
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
	return fmt.Sprintf("%.2f", value)
}
