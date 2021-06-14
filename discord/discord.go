package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ta93-ito/discord-weather-bot/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func DiscordNew() {
	discord, err := discordgo.New()
	discord.Token = config.Config.Token

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discord.AddHandler(onMessageCreate)

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
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//botName := config.Config.BotName

	_, err := discordgo.New()

	if err != nil {
		log.Println("Error getting channel: ", err)
		return

		fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
	}
}