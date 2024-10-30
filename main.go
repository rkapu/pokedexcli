package main

import (
	"time"

	"github.com/rkapu/pokedexcli/internal/pokeapi"
)

const prefix = "pokedex > "

func main() {
	cfg := &config{
		pokeapiClient:            pokeapi.NewClient(time.Hour),
		locationAreasNextUrl:     nil,
		locationAreasPreviousUrl: nil,
		caughtPokemon:            map[string]pokeapi.PokemonResp{},
	}

	startREPL(cfg)
}
