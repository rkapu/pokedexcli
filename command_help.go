package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Available commands:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s - %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}
