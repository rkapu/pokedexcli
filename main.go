package main

import "github.com/rkapu/pokedexcli/internal/pokeapi"

const prefix = "pokedex > "

type Config struct {
	pokeapiClient            pokeapi.Client
	LocationAreasNextUrl     *string
	LocationAreasPreviousUrl *string
}

func main() {
	cfg := Config{
		pokeapiClient:            pokeapi.NewClient(),
		LocationAreasNextUrl:     nil,
		LocationAreasPreviousUrl: nil,
	}
	startREPL(&cfg)
}
