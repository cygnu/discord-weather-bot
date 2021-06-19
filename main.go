package main

import (
	"github.com/ta93-ito/discord-weather-bot/apis/discord"
	"github.com/ta93-ito/discord-weather-bot/server"
)

func main() {
	discord.DiscordNew()
	server.StartWebServer()
}
