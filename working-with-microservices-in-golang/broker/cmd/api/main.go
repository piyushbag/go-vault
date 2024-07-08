package main

import (
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting broker service on port %s\n", webPort)

	// define the server
	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	// start the server
	if err := srv.ListenAndServe(); err != nil {
		log.Panicf("server failed to start: %v", err)
	}

}
