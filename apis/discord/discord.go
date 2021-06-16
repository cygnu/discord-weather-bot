package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ta93-ito/discord-weather-bot/apis/openweather"
	_ "github.com/ta93-ito/discord-weather-bot/apis/openweather"
	"github.com/ta93-ito/discord-weather-bot/config"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func DiscordNew() {
	discord, err := discordgo.New()
	discord.Token = config.Config.Token

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discord.AddHandler(messageCreate)
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer discord.Close()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-ch
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, "/") {
		return
	}
	city := strings.Replace(m.Content, "/", "", 1)

	weather := openweather.GetCurrentWeather(city)
	s.ChannelMessageSend(m.ChannelID, weather.Name)
}
