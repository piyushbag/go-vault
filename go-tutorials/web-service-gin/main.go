/*
// 1. api design
/albums GET POST
/albums/:id GET
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// api structure
type user struct {
	ID    int
	Name  string
	Email string
}

// create data to be sourced as api
var users = []user{
	{ID: 1, Name: "Tom", Email: "tom@gmail.com"},
	{ID: 2, Name: "Dick", Email: "dick@gmail.com"},
	{ID: 3, Name: "Harry", Email: "harry@gmail.com"},
}

// marshal and send your data
func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// unmarshal and add the record in your memory
func createUser(c *gin.Context) {
	var newUser user
	err := c.BindJSON(&newUser)
	if err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusAccepted, newUser)
}

// get a single id
func getUserById(c *gin.Context) {
	id := c.Param("id")
	int_id, _ := strconv.Atoi(id)
	for _, user := range users {
		if user.ID == int_id {
			c.IndentedJSON(http.StatusOK, id)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
	router := gin.Default()

	router.GET("/albums", getUser)
	router.POST("/albums", createUser)
	router.GET("/albums/:id", getUserById)

	router.Run("localhost:8080")

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
