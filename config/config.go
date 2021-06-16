package config

import (
	"gopkg.in/go-ini/ini.v1"
	"log"
)

type ConfigList struct {
	Token   string
	BotName string
	ApiKey  string
	Port    int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	Config = ConfigList{
		Token:   cfg.Section("discord_bot").Key("token").String(),
		BotName: cfg.Section("discord_bot").Key("bot_name").String(),
		ApiKey:  cfg.Section("openweather").Key("api_key").String(),
		Port:    cfg.Section("web").Key("port").MustInt(),
	}
}
