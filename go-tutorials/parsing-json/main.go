package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Name    string
	Pokemon []Pokemon
}

type Pokemon struct {
	entry   int
	species PokemonSpecies
}

type PokemonSpecies struct {
	Name string
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		log.Println("Could not fetch data from the given api")
		return
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("could not parse the data")
		return
	}

	fmt.Println(string(data))

	var responseObject Response
	json.Unmarshal(data, &responseObject)

	fmt.Println("test", responseObject)

	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].species.Name)
	}
}
