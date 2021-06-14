package openweather

import (
	"encoding/json"
	"fmt"
	"github.com/ta93-ito/discord-weather-bot/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

const Endpoint = "https://api.openweathermap.org/data/2.5/weather"

func GetCurrentWeather(city string) string {
	token := config.Config.ApiKey
	//city := "Osaka,jp"

	values := url.Values{}
	values.Set("q", city)
	values.Set("appid", token)

	res, err := http.Get(fmt.Sprintf("%s?%s", Endpoint, values.Encode()))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var apiRes OpenWeather
	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}

	// fmt.Printf("場所: %v\n", city)
	// fmt.Printf("時間: %s\n", time.Unix(int64(apiRes.Dt), 0))
	// fmt.Printf("天気: %s\n", apiRes.Weather[0].Main)
	return apiRes.Weather[0].Main
}

type OpenWeather struct {
	Coord      Coord   `json:"coord"`
	Weather    Weather `json:"weather"`
	Base       string  `json:"base"`
	Main       Main    `json:"main"`
	Visibility int     `json:"visibility"`
	Wind       Wind    `json:"wind"`
	Clouds     Clouds  `json:"clouds"`
	Dt         int     `json:"dt"`
	Sys        Sys     `json:"sys"`
	Timezone   int     `json:"timezone"`
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Cod        int     `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather []struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int     `json:"type"`
	Id      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}
