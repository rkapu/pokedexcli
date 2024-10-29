package main

import (
	"fmt"
	"github.com/rkapu/pokedexcli/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.GetNextLocationAreas()

	if err != nil {
		return err
	}

	for _, locationArea := range resp.Results {
		fmt.Println(locationArea.Name)
	}

	fmt.Println()

	return nil
}
