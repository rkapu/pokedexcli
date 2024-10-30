package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.locationAreasNextUrl)

	if err != nil {
		return err
	}

	cfg.locationAreasNextUrl = resp.Next
	cfg.locationAreasPreviousUrl = resp.Previous

	for _, locationArea := range resp.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.locationAreasPreviousUrl == nil {
		return fmt.Errorf("Nowhere to go back")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.locationAreasPreviousUrl)

	if err != nil {
		return err
	}

	cfg.locationAreasNextUrl = resp.Next
	cfg.locationAreasPreviousUrl = resp.Previous

	for _, locationArea := range resp.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
