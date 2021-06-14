package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ta93-ito/discord-weather-bot/config"
	"github.com/ta93-ito/discord-weather-bot/openweather"
	"os"
	"os/signal"
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
	signal.Notify(ch, os.Interrupt)

	go func() {
		for c := range ch {
			fmt.Println(c)
			close(ch)
			os.Exit(1)
		}
	}()

	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	weather := openweather.GetCurrentWeather(m.Content)
	s.ChannelMessageSend(m.ChannelID, weather)
}
