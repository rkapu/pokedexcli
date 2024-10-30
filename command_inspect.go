package main

import (
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, exist := cfg.caughtPokemon[pokemonName]
	if !exist {
		return fmt.Errorf("You have not caught %s yet", pokemonName)
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, tp := range pokemon.Types {
		fmt.Printf(" - %s\n", tp.Type.Name)
	}

	return nil
}
