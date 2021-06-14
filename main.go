package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Endpoint = "api.openweathermap.org/data/2.5/weather"

func main() {
	handler()
}

func handler() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.Run(":8080")
}
