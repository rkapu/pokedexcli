package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL(cfg *Config) {
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
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg)
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

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
			description: "Prints previous page of area locations",
			callback:    callbackMapb,
		},
	}
}
