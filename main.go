package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func commandMap() error {
	areas, err := GetLocationAreas(locationAreasUrl)
	if IsErr(err) {
		fmt.Printf("Error! %v", err)
		return err
	}
	fmt.Println(areas[0].Results[0].Name)
	return nil
}

func commandMapB() error {
	return nil
}

func commandHealth() error {
	if ok, status := TestApi(); ok {
		fmt.Println("Connection and response from API is good!")
	} else {
		fmt.Printf("Failed to receive good response from API; status %v\n", status)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			config:      nil,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      nil,
		},
		"map": {
			name:        "map",
			description: "Shows a list of 20 location areas in the Pokemon world (call Map again for the next 20)",
			callback:    commandMap,
			config:      nil,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 location areas after calling map",
			callback:    commandMapB,
			config:      nil,
		},
		"health": {
			name:        "health",
			description: "Checks the API and connection health",
			callback:    commandHealth,
			config:      nil,
		},
	}
}
