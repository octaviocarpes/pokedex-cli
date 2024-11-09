package poke_api

import (
	http_client "github.com/octaviocarpes/pokedex-cli/http-client"
	cache "github.com/octaviocarpes/pokedex-cli/poke-api/cache"
	pokemon "github.com/octaviocarpes/pokedex-cli/pokemon"
)

func GetPokemon(name string) (pokemon.Pokemon, error) {
	url := "/pokemon/" + name

	cachedPokemon, cachedPokemonError := cache.GetCachedData[pokemon.Pokemon](url)

	if cachedPokemonError == nil {
		return cachedPokemon, nil
	}

	response, err := http_client.ExecuteGet[pokemon.Pokemon](url)

	return response, err
}
