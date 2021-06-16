package main

import (
	"github.com/ta93-ito/discord-weather-bot/apis/discord"
	"github.com/ta93-ito/discord-weather-bot/server"
)

var Endpoint = "api.openweathermap.org/data/2.5/weather"

func main() {
	discord.DiscordNew()
	server.StartWebServer()
}
