package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TestApi() (bool, int) {
	// makes a GET call to languages to confirm a 200 response is received

	res, err := http.Get("https://pokeapi.co/api/v2/language/1/")
	if IsErr(err) {
		return false, -1
	}

	if res.StatusCode != http.StatusOK {
		return false, res.StatusCode
	}

	return true, http.StatusOK

}

func GetLocationAreas(apiPath string) (locationAreaResponse, error) {

	var blankResponse locationAreaResponse

	res, err := http.Get(apiPath)
	if IsErr(err) {
		return blankResponse, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return blankResponse, fmt.Errorf("unexpected status: %s", res.Status)
	}

	var result locationAreaResponse

	if err := json.NewDecoder(res.Body).Decode(&result); IsErr(err) {
		return blankResponse, fmt.Errorf("failed to decode: %w", err)
	}

	return result, nil

}
