# weather-go

## Usage

### Help

Running `$ ./widget -h` results in

```shell
Usage of ./widget:
  -env string
    	Path to .env file containing api key and optional location (default "./.env")
  -location string
    	Set location to get weather (default "New York")
  -mode string
    	Set output mode
    	small: Print single line consisting of location name, temperature, weather
    	full: Prints full data of current weather (default "small")
```

### Use as Widget

Running `$ ./widget -location rome -env <path to env file> -mode small` results in

```shell
Rome: 10.37°C light rain
```

### Full Weather Data

Running `$ ./widget -location rome -env <path to env file> -mode full` results in

```shell
Rome:
	Rain, light rain
	Temperature: 10.37 °C
	Humidity: 78 hpa
	Pressure: 1007
	Wind: 1.540000 m/s from 220
```

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

### OpenWeatherMap

OpenWeatherMap provides an API to geocode the given location to latitude and longitude coordinates. In addition to that it provides an API to get the current weather of a location and the forecast.

Each request takes one ticket. Calling the script currently takes two credits, one for geocoding the location, and one to get the data.
