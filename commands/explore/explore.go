package explore_command

import (
	"fmt"
	"log"

	api "github.com/octaviocarpes/pokedex-cli/poke-api"
)

func ListEncounters(location string) error {

	response, err := api.GetLocationArea(location)

	if err != nil {
		log.Fatal("failed to explore")
	}

	for _, encounter := range response.PokemonEncounters {
		fmt.Printf("%v\n", encounter.Pokemon.Name)
	}

	return nil
}
