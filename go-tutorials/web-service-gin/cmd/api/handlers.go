package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums returns a list of albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	Albums = append(Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album in the slice by its ID and returns it.
// parameter sent by the client, then returns the album as JSON.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
