package commands

import (
	"fmt"
	"os"

	map_command "github.com/octaviocarpes/pokedex-cli/commands/map"
)

type cliCommand struct {
	name     string
	Callback func() error
}

func HelpCallback() error {
	fmt.Println(`
Here is a list of all available cli Commands:

	- help: Shows how to use the Pokedex CLI

	- exit: Exits the Program

	- map: Fetch the locations of the Pokemon World (Onwards)
	
	- mapb: Fetch the locations of the Pokemon World (Backwards)
	`)

	return nil
}

func exitCallback() error {
	os.Exit(0)
	return nil
}

func mapCallback() error {
	err := map_command.GetLocations()
	return err
}

func mapbCallback() error {
	err := map_command.GetBackwardLocations()
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
	}

	return AllCommands
}
