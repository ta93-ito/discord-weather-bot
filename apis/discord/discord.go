package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ta93-ito/discord-weather-bot/apis/openweather"
	"github.com/ta93-ito/discord-weather-bot/config"
	"os"
	"os/signal"
	"strings"
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

	fmt.Printf("%s %s %s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, city)

	res, err := openweather.GetForecast(city)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	s.ChannelMessageSend(m.ChannelID, SyntheticMessage(res.Forecasts, city))
}

func SyntheticMessage(list []openweather.Forecast, city string) string {
	var each_weather []string
	necessary_list := list[3:7]
	for i := 0; i < len(necessary_list); i++ {
		fomatted_dt := strings.Replace(necessary_list[i].DtTxt[5:10], "-", "月", -1) + "日" + " " + necessary_list[i].DtTxt[11:13] + "時"
		each_weather = append(each_weather, fomatted_dt + " " + necessary_list[i].Weather[0].Description)
	}

	msg := city + "の天気\n" + strings.Join(each_weather, "\n")
	return msg
}