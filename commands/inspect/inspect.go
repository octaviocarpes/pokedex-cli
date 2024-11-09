package inspect_command

import pokedex "github.com/octaviocarpes/pokedex-cli/pokedex"

func Inspect(name string) error {
	pokedex.InspectPokemon(name)
	return nil
}
