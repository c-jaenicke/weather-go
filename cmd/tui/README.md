# weather-go - TUI

A small TUI version of the widget version, displays current weather and 24 hour forecast.

Built using [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea)
and [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss).

## Screenshot

![screenshot](https://raw.githubusercontent.com/c-jaenicke/weather-go/main/images/tui.png)

## Usage

Call the application using `./tui -env <path to .env file>`.

Enter your location in the input box and confirm using the enter key.

## .env File

The `.env`-file has to at least contain the `API_KEY` value, mapped to your OpenWeatherMap API Key!
The `location` value is optional, in case you don't want to call the script with the location. When doing so, the script needs to be called with `.env` as the location parameter!

```.env
# Example .env file
# OpenWeatherMap API Key
API_KEY=<OpenWeatherMap API Key>
# Optional
LOCATION=<city_name,state_code,country_code>
```
