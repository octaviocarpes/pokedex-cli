package poke_api

import (
	"log"
	"strconv"

	http_client "github.com/octaviocarpes/pokedex-cli/http-client"
	cache "github.com/octaviocarpes/pokedex-cli/poke-api/cache"
	utils "github.com/octaviocarpes/pokedex-cli/utils"
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

func cacheLocation(url string, location ListLocationsResponse) {
	byteSlice := utils.ConvertToByteSlice(location)

	if byteSlice != nil {
		cache.Add(url, byteSlice)
	}
}

func getCachedLocation(url string) (ListLocationsResponse, error) {
	cachedResponse, ok := cache.Get(url)
	response := &ListLocationsResponse{}

	if ok {
		parsedResponse, conversionError := utils.ConvertFromByteSlice[ListLocationsResponse](cachedResponse)

		if conversionError != nil {
			log.Fatal("Failed to decode data from byte slice")
		}

		return parsedResponse, nil
	}

	return *response, nil
}

func GetLocations(offset int) (ListLocationsResponse, error) {
	stringValue := strconv.Itoa(offset)
	search := "offset=" + stringValue
	url := "/location?" + search

	cachedLocation, getCachedError := getCachedLocation(url)

	if getCachedError != nil {
		log.Fatal(getCachedError)
	}

	if cachedLocation.Results != nil {
		return cachedLocation, nil
	}

	response, error := http_client.ExecuteGet[ListLocationsResponse](url)

	cacheLocation(url, response)

	if error != nil {
		return response, error
	}

	return response, nil
}
