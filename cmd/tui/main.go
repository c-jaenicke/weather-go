package main

import (
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"weather-go/pkg/env"
	"weather-go/pkg/weather"
)

func main() {
	// gets flags
	var pathEnv string
	var location string
	flag.StringVar(&pathEnv, "env", "", "Path to .env file containing api key and optional location")
	flag.StringVar(&location, "location", "", "Set location to get weather.\nExpected format: <city_name,state_code,country_code>\nCan use '.env' as value, doing so will load the location from the given .env file")
	flag.Parse()

	if pathEnv == "" {
		log.Fatal("No .env file specified! Exiting...")
	}

	env.Path = pathEnv

	// get location from .env file if ".env" given as location
	if location == ".env" {
		var err error
		location, err = env.GetEnv("LOCATION")
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	// start bubbletea
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func currentWeather(m Model) (string, string, error) {
	weatherData, locationName, err := weather.GetWeatherData(m.location)
	if err != nil {
		log.Println("currentWeather: failed to get current weather: " + err.Error())
		return "", m.location, err
	}
	return fmt.Sprintf("%s, %s"+
			"\nTemperature: %s °C"+
			"\nHumidity: %d hpa"+
			"\nPressure: %d"+
			"\nWind: %s m/s from %d",
			weatherData.Current.Weather[0].Main,
			weatherData.Current.Weather[0].Description,
			shortenFloat(weatherData.Current.Temp),
			weatherData.Current.Humidity,
			weatherData.Current.Pressure,
			shortenFloat(weatherData.Current.WindSpeed),
			weatherData.Current.WindDeg),
		locationName, nil
}

// getForecast returns forecast as table.Model.View() string
func getForecast(m Model) (string, error) {
	forecastObject, _, err := weather.GetForecast(m.location)
	if err != nil {
		log.Println("getForecast: failed to get forecast: " + err.Error())
		return "", err
	}

	var sb strings.Builder

	fmt.Println(len(forecastObject.List))

	for index, i := range forecastObject.List {
		if index == 5 {
			break
		}
		text := fmt.Sprintf("\n%s"+
			"\n%s"+
			"\n\tTemp: %s °C"+
			"\n\tHumidity: %d%%"+
			"\n\tPressure: %d hpa"+
			"\n\tWind: %s m/s from %d"+
			"\n\n", unixToDateTime(i.Dt), i.Weather[0].Description, shortenFloat(i.Main.Temp), i.Main.Humidity, i.Main.Pressure, shortenFloat(i.Wind.Speed), i.Wind.Deg)

		sb.WriteString(text)
	}

	return sb.String(), nil
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
