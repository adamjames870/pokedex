package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {

	commands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println("\nExiting...")
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		words := cleanInput(input)
		firstWord := words[0]

		cmd, exists := commands[firstWord]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
