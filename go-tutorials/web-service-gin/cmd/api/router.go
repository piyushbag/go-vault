package main

import "github.com/gin-gonic/gin"

func (app *Config) routes() {
	router := gin.Default()
	router.GET("/albums", getAlbums)        // curl http://localhost:8080/albums
	router.GET("/albums/:id", getAlbumByID) // curl http://localhost:8080/albums/1
	router.POST("/albums", postAlbums)      // curl -X POST http://localhost:8080/albums -d '{"id": "4", "title": "The Modern Sound of Betty Carter", "artist": "Betty Carter", "price": "49.99"}' -H "Content-Type: application/json"
}
