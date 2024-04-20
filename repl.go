package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eneax/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		reader.Scan()
		text := reader.Text()

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		availableCommands := getCommands()

		command, exists := availableCommands[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(input string) []string {
	loweredInput := strings.ToLower(input)
	words := strings.Fields(loweredInput)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List next location areas",
			callback:    commandMapNext,
		},
		"map-back": {
			name:        "map-back",
			description: "List previous location areas",
			callback:    commandMapPrev,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commands
}
