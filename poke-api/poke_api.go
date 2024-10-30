package poke_api

import (
	"strconv"

	http_client "github.com/octaviocarpes/pokedex-cli/http-client"
	cache "github.com/octaviocarpes/pokedex-cli/poke-api/cache"
)

type ListLocationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(offset int) (ListLocationsResponse, error) {
	stringValue := strconv.Itoa(offset)
	search := "offset=" + stringValue
	url := "/location?" + search

	cachedResponse := cache.Get(url)

	if res, ok := cachedResponse.(ListLocationsResponse); ok {
		return res, nil
	}

	response, error := http_client.ExecuteGet[ListLocationsResponse](url)

	if error != nil {
		return response, error
	}

	return response, nil
}
