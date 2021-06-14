package config

import (
	"gopkg.in/go-ini/ini.v1"
	"log"
)

type ConfigList struct {
	ApiKey string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file\n", err)
	}

	Config = ConfigList{
		ApiKey: cfg.Section("openweather").Key("api_key").String(),
	}
}
