package main

import (
	"fmt"
)

func callbackMap(cfg *Config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.LocationAreasNextUrl)

	if err != nil {
		return err
	}

	for _, locationArea := range resp.Results {
		fmt.Println(locationArea.Name)
	}

	fmt.Println()

	cfg.LocationAreasNextUrl = resp.Next
	cfg.LocationAreasPreviousUrl = resp.Previous

	return nil
}

func callbackMapb(cfg *Config) error {
	if cfg.LocationAreasPreviousUrl == nil {
		return fmt.Errorf("Nowhere to go back")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.LocationAreasPreviousUrl)

	if err != nil {
		return err
	}

	for _, locationArea := range resp.Results {
		fmt.Println(locationArea.Name)
	}

	fmt.Println()

	cfg.LocationAreasNextUrl = resp.Next
	cfg.LocationAreasPreviousUrl = resp.Previous

	return nil
}
