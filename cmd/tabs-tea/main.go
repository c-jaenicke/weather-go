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
	var pathEnv string
	flag.StringVar(&pathEnv, "env", "./.env", "Path to .env file containing api key and optional location")
	flag.Parse()

	env.Path = pathEnv

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// currentWeather gets current weather from a location, returns as formatted string
func currentWeather(m Model) string {
	weatherData, locationName := weather.GetWeatherData(m.location)
	return fmt.Sprintf("%s: \n\t%s, %s\n\tTemperature: %s °C\n\tHumidity: %d hpa\n\tPressure: %d\n\tWind: %f m/s from %d", locationName, weatherData.Current.Weather[0].Main, weatherData.Current.Weather[0].Description, strconv.FormatFloat(weatherData.Current.Temp, 'f', -1, 64), weatherData.Current.Humidity, weatherData.Current.Pressure, weatherData.Current.WindSpeed, weatherData.Current.WindDeg)
}

func getForecast(m Model) string {
	forecastObject, _ := weather.GetForecast(m.location)

	var sb strings.Builder

	fmt.Println(len(forecastObject.List))

	for _, i := range forecastObject.List {
		text := fmt.Sprintf("\n%s"+
			"\n%s"+
			"\n\tTemp: %s °C"+
			"\n\tHumidity: %d%%"+
			"\n\tPressure: %d hpa"+
			"\n\tWind: %s m/s from %d"+
			"\n\n", unixToDateTime(i.Dt), i.Weather[0].Description, shortenFloat(i.Main.Temp), i.Main.Humidity, i.Main.Pressure, shortenFloat(i.Wind.Speed), i.Wind.Deg)

		sb.WriteString(text)
	}

	return sb.String()
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
