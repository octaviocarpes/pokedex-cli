package poke_api

import (
	"log"

	http_client "github.com/octaviocarpes/pokedex-cli/http-client"
	cache "github.com/octaviocarpes/pokedex-cli/poke-api/cache"
)

type LocationAreaList struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationArea(search string) (LocationAreaList, error) {
	url := "/location-area/" + search

	cachedResponse, getCachedError := cache.GetCachedData[LocationAreaList](url)

	if getCachedError != nil {
		log.Fatal(getCachedError)
	}

	if cachedResponse.EncounterMethodRates != nil {
		return cachedResponse, nil
	}

	response, error := http_client.ExecuteGet[LocationAreaList](url)

	cache.SaveDataInCache(url, response)

	if error != nil {
		return response, error
	}

	return response, nil
}
