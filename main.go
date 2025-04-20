package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}
