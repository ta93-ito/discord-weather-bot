package main

import (
	"github.com/ta93-ito/discord-weather-bot/discord"
	"github.com/ta93-ito/discord-weather-bot/openweather"
)
var Endpoint = "api.openweathermap.org/data/2.5/weather"

func main() {
	openweather.GetCurrentWeather("Tokyo,jp")
	discord.DiscordNew()
	StartWebServer()
}
