package main

import (
	"github.com/ta93-ito/notify-absentee/geo"
	"net/http"

	"github.com/gin-gonic/gin"
)

const Endpoint = "api.openweathermap.org/data/2.5/weather"

func main() {
	geo.GetGeoJson()
	handler()
}

func handler() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.Run(":8080")
}
