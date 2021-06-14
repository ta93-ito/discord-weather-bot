package geo

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antonholmquist/jason"
)


const url = "https://get.geojs.io/v1/ip/geo.json"

func GetGeoJson()  {
	res, _ := http.Get(url)
	defer res.Body.Close()

	byteArray, _ := ioutil.ReadAll(res.Body)
	str := string(byteArray)

	v, err := jason.NewObjectFromBytes([]byte(str))
	if err != nil{
		panic(err)
	}

	latitude, _ := v.GetString("latitude")
	longitude, _ := v.GetString("longitude")

	fmt.Println("latitude:" + latitude)
	fmt.Println("longitude:" + longitude)
}
