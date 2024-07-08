package main

import (
	"context"
	"log"
	"logger/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRPCPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect to the MongoDB database
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	// create context for the server
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection to MongoDB
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// start the server
	// go app.server()

	// create a new server
	log.Printf("Starting logger service on port %s\n", webPort)
	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	// start the server
	if err = srv.ListenAndServe(); err != nil {
		log.Panicf("server failed to start: %v", err)
	}
}

// func (c *Config) server() {
// 	// create a new server
// 	srv := &http.Server{
// 		Addr:    ":" + webPort,
// 		Handler: c.routes(),
// 	}

// 	// start the server
// 	if err := srv.ListenAndServe(); err != nil {
// 		log.Panicf("server failed to start: %v", err)
// 	}
// }

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}
