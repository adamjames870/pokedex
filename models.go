package main

import "time"

const locationAreasUrl string = "https://pokeapi.co/api/v2/location-area/"
const defaultDuration time.Duration = 5 * time.Second

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous *string        `json:"previous"` // nullable field
	Results  []locationArea `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*commandConfig) error
	config      *commandConfig
}

type commandConfig struct {
	nextUrl string
	prevUrl string
}
