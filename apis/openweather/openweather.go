package openweather

import (
	"encoding/json"
	"fmt"
	"github.com/ta93-ito/discord-weather-bot/apis/geocoding"
	"github.com/ta93-ito/discord-weather-bot/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

const Endpoint1 = "https://api.openweathermap.org/data/2.5/weather"
const Endpoint2 = "https://api.openweathermap.org/data/2.5/forecast"

func GetCurrentWeather(city string) string {
	lat, lon := geocoding.Geocoding(city)

	token := config.Config.ApiKey

	values := url.Values{}
	values.Set("lat", lat)
	values.Set("lon", lon)
	values.Set("appid", token)

	res, err := http.Get(fmt.Sprintf("%s?%s", Endpoint1, values.Encode()))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	switch res.Status[0:1] {
	case "4":
		return "invalid statement!"
	case "5":
		return "something went wrong..."
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var apiRes OpenWeather

	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}
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

func GetForcast(city string) string {
	lat, lon := geocoding.Geocoding(city)

	token := config.Config.ApiKey

	values := url.Values{}
	values.Set("lat", lat)
	values.Set("lon", lon)
	values.Set("appid", token)

	res, err := http.Get(fmt.Sprintf("%s?%s", Endpoint2, values.Encode()))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var apiRes Forecast

	if err := json.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}
	return apiRes[0].DtTxt
}

type Forecast []struct {
	Main2 Main2  `json:"main"`
	DtTxt string `json:"dt_txt"`
}

type Main2 struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}
