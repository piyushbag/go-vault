package main

import (
	"log"
	"net/http"
)

type Config struct{}

const webPort = ":80"

func main() {
	app := Config{}
	log.Println("Starting server on", webPort)

	srv := &http.Server{
		Addr:    webPort,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println("server error:", err)
	}
}
