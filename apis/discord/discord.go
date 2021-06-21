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
	if err != nil {
		fmt.Println(err)
	}

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

	return
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
	var eachWeather []string
	necessaryList := list[3:7]

	for i := 0; i < len(necessaryList); i++ {
		formattedDt := FormatDtTxt(necessaryList[i].DtTxt)
		eachWeather = append(eachWeather, fmt.Sprintf("%s %s", formattedDt, necessaryList[i].Weather[0].Description))
	}

	msg := fmt.Sprintf("%sの天気\n%s\n", city, strings.Join(eachWeather, "\n"))
	return msg
}

func FormatDtTxt(dt_txt string) string {
	split_date := dt_txt[5:10]
	split_time := dt_txt[11:13]
	date := &split_date
	time := &split_time

	if strings.HasPrefix(*date, "0") {
		 *date = strings.Replace(*date, "0", "", 1)
	}
	if strings.HasPrefix(*time, "0") {
		*time = strings.Replace(*time, "0", "", 1)
	}
	return fmt.Sprintf("%s日 %s時", strings.Replace(*date, "-", "月", 1), *time)
}
