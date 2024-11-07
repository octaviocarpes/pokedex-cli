package explore_command

import (
	"fmt"

	api "github.com/octaviocarpes/pokedex-cli/poke-api"
)

func ListEncounters(location string) error {

	response, err := api.GetLocationArea(location)

	if err != nil {
		return err
	}

	for _, encounter := range response.PokemonEncounters {
		fmt.Printf("%v\n", encounter.Pokemon.Name)
	}

	return nil
}
