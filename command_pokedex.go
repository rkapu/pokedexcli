package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your pokedex is empty")
		return nil
	}

	fmt.Println("Your pokedex:")
	for pokemonName := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}
