package commands

import (
	"fmt"
	"os"

	catch_command "github.com/octaviocarpes/pokedex-cli/commands/catch"
	explore_command "github.com/octaviocarpes/pokedex-cli/commands/explore"
	map_command "github.com/octaviocarpes/pokedex-cli/commands/map"
)

type cliCommand struct {
	name     string
	Callback func(args []string) error
}

func HelpCallback(args []string) error {
	fmt.Println(`
Here is a list of all available cli Commands:

	- help: Shows how to use the Pokedex CLI

	- exit: Exits the Program

	- map: Fetch the locations of the Pokemon World (Onwards)
	
	- mapb: Fetch the locations of the Pokemon World (Backwards)

	- explore: Explore a location of the Pokemon world and shows a list of Pokemons you can encounter.
		To use the explore command you need to pass a location as a parameter eg:
			   
		Pokedex CLI > explore mt-coronet-2f
	
	- catch: Attempt to catch a Pokemon
		To use the catch command you need to pass a pokemon as a parameter eg:
			   
		Pokedex CLI > catch pikachu
	`)

	return nil
}

func exitCallback(args []string) error {
	os.Exit(0)
	return nil
}

func mapCallback(args []string) error {
	err := map_command.GetLocations()
	return err
}

func mapbCallback(args []string) error {
	err := map_command.GetBackwardLocations()
	return err
}

func exploreCallback(args []string) error {
	err := explore_command.ListEncounters(args[0])
	return err
}

func catchCallback(args []string) error {
	err := catch_command.Catch(args[0])
	return err
}

func ListAllCommands() map[string]cliCommand {
	AllCommands := map[string]cliCommand{
		"help": {
			name:     "help",
			Callback: HelpCallback,
		},
		"exit": {
			name:     "exit",
			Callback: exitCallback,
		},
		"map": {
			name:     "map",
			Callback: mapCallback,
		},
		"mapb": {
			name:     "mapb",
			Callback: mapbCallback,
		},
		"explore": {
			name:     "explore",
			Callback: exploreCallback,
		},
		"catch": {
			name:     "catch",
			Callback: catchCallback,
		},
	}

	return AllCommands
}
