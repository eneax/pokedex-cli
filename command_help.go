package main

import "fmt"

func callbackHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Available commands:")
	fmt.Println()

	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}

	fmt.Println("")
	return nil
}
