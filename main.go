package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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

		fmt.Println("Your command was:", firstWord)

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
