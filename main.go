package main

import (
	"time"

	"github.com/rkapu/pokedexcli/internal/pokeapi"
)

const prefix = "pokedex > "

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startREPL(cfg)
}
