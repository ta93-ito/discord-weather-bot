package main

import (
	"github.com/ta93-ito/notify-absentee/openweather"
)

var Endpoint = "api.openweathermap.org/data/2.5/weather"

func main() {
	openweather.GetCurrentWeather("Tokyo,jp")
	StartWebServer()
}
