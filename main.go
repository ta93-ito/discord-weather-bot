package main

import (
	"github.com/ta93-ito/discord-weather-bot/discord"
)
var Endpoint = "api.openweathermap.org/data/2.5/weather"

func main() {
	discord.DiscordNew()
	StartWebServer()
}
