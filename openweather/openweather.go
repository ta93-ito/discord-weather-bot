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

	fmt.Println(apiRes)
	return apiRes.Weather[0].Main
}

type OpenWeather struct {
	Weather  Weather `json:"weather"`
	Main     Main    `json:"main"`
	Timezone int     `json:"timezone"`
	Name     string  `json:"name"`
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
