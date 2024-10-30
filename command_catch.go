package main

import (
	"fmt"
	"math/rand"
)

const catchThreshold = 30

func callbackCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	maxRand := max(pokemon.BaseExperience, catchThreshold)
	rand := rand.Intn(maxRand)

	if rand > catchThreshold {
		fmt.Printf("%s escaped!\n\n", pokemonName)
		return nil
	}

	cfg.caughtPokemon[pokemonName] = pokemon

	fmt.Printf("%s was caught!\n\n", pokemonName)

	return nil
}
