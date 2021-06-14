package main

import (
	"github.com/ta93-ito/notify-absentee/discord"
	"github.com/ta93-ito/notify-absentee/openweather"
)

var Endpoint = "api.openweathermap.org/data/2.5/weather"

func main() {
	openweather.GetCurrentWeather("Tokyo,jp")
	discord.DiscordNew()
	StartWebServer()
}
