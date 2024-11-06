package poke_api

import (
	"log"
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
	url := "/location-area?" + search

	cachedLocation, getCachedError := cache.GetCachedData[ListLocationsResponse](url)

	if getCachedError != nil {
		log.Fatal(getCachedError)
	}

	if cachedLocation.Results != nil {
		return cachedLocation, nil
	}

	response, error := http_client.ExecuteGet[ListLocationsResponse](url)

	cache.SaveDataInCache(url, response)

	if error != nil {
		return response, error
	}

	return response, nil
}
