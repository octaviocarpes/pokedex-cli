package pokedex

import (
	"fmt"
	"maps"
	"slices"

	pokemon "github.com/octaviocarpes/pokedex-cli/pokemon"
)

var pokedex = make(map[string]pokemon.Pokemon)

func printPokemon(pokemon pokemon.Pokemon) {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")

	for _, stat := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")

	for _, pType := range pokemon.Types {
		fmt.Printf("-%v\n", pType.Type.Name)
	}
}

func AddToPokedex(name string, pokemon pokemon.Pokemon) {
	pokedex[pokemon.Name] = pokemon
}

func CheckPokedex() {
	fmt.Printf("Your Pokedex:\n")

	pokemons := slices.Collect(maps.Keys(pokedex))

	if len(pokemons) == 0 {
		fmt.Printf("Your pokedex is empty\n")
		return
	}

	for _, pokemon := range pokemons {
		fmt.Printf(" - %v\n", pokemon)
	}

	fmt.Println()
}

func InspectPokemon(name string) {
	pokemon, ok := pokedex[name]

	if !ok {
		fmt.Printf("You haven't caught this Pokemon yet\n\n")
		return
	}

	printPokemon(pokemon)
}
