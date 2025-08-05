package main

const locationAreasUrl string = "https://pokeapi.co/api/v2/location-area/"

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *commandConfig
}

type commandConfig struct {
	nextUrl string
	prevUrl string
}
