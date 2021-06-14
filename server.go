package main

import (
	"fmt"
	"github.com/ta93-ito/notify-absentee/config"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func StartWebServer() error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
