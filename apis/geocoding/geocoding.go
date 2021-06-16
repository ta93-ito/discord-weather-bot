package geocoding

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const Geocoding_Endpoint = "https://www.geocoding.jp/api"

func Geocoding(city string) (string, string) {
	values := url.Values{}
	values.Set("q", city)

	res, _ :=  http.Get(fmt.Sprintf("%s?%s", Geocoding_Endpoint, values.Encode()))
	bytes, _ := ioutil.ReadAll(res.Body)

	var apiRes Geocode

	if err := xml.Unmarshal(bytes, &apiRes); err != nil {
		panic(err)
	}
	return apiRes.Coordinate.Lat, apiRes.Coordinate.Lng
}

type Geocode struct {
	Address    string     `xml:"address"`
	Coordinate Coordinate `xml:"coordinate"`
}

type Coordinate struct {
	Lat     string `xml:"lat"`
	Lng     string `xml:"lng"`
}
