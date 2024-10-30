package main

import "fmt"

func callbackExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No location name provided for exploration")
	}
	locationAreaName := args[0]

	fmt.Printf("Exploring %s...\n", locationAreaName)
	resp, err := cfg.pokeapiClient.ShowLocationArea(locationAreaName)

	if err != nil {
		return err
	}

	if len(resp.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found")
		return nil
	}

	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range resp.PokemonEncounters {
		fmt.Println(" - ", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
