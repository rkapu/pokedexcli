package main

import "github.com/rkapu/pokedexcli/internal/pokeapi"

const prefix = "pokedex > "

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startREPL(cfg)
}
