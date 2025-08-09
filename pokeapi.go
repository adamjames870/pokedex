package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PokeApi struct {
	pk *pokeCache
}

func NewPokeApi() PokeApi {
	cache := NewCache(5 * time.Second)
	pa := &PokeApi{
		pk: &cache,
	}
	pa.pk.add("Initial", nil)
	return *pa
}

func (a *PokeApi) TestApi() (bool, int) {
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

func (a *PokeApi) GetLocationAreas(apiPath string) (locationAreaResponse, error) {

	cacheres, exists := a.pk.Get(apiPath)
	if exists {
		var response locationAreaResponse
		json.Unmarshal(cacheres, &response)
		a.pk.Renew(apiPath)
		return response, nil
	} else {
		apires, err := getLocationAreasFromApi(apiPath)
		if IsErr(err) {
			return apires, err
		}
		cacheval, _ := json.Marshal(apires)
		a.pk.add(apiPath, cacheval)
		return apires, nil
	}

}

func getLocationAreasFromApi(apiPath string) (locationAreaResponse, error) {
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
