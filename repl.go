package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rkapu/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient            pokeapi.Client
	locationAreasNextUrl     *string
	locationAreasPreviousUrl *string
	caughtPokemon            map[string]pokeapi.PokemonResp
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	availableCommands := getCommands()
	for {
		fmt.Print(prefix)
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	trimmed := strings.TrimSpace(lowered)
	words := strings.Fields(trimmed)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    callbackExit,
		},
		"help": {
			name:        "help",
			description: "Prints a help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Lists next page of area locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous page of area locations",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try to catch a pokemon",
			callback:    callbackCatch,
		},
	}
}
