package commands

import (
	"fmt"
	"log"
	"os"

	poke_api "github.com/octaviocarpes/pokedex-cli/poke-api"
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

	- map: Fetch the locations of the Pokemon World
	`)

	return nil
}

func exitCallback() error {
	os.Exit(0)
	return nil
}

func mapCallback() error {
	// TODO: create map mode logic
	response, error := poke_api.GetLocations(0)

	if error != nil {
		log.Fatal("Failed to fetch pokeapi locations")
		os.Exit(1)
	}

	fmt.Printf("Pokemon Locations: \n\n")

	for _, locataion := range response.Results {
		fmt.Printf("%s\n", locataion.Name)
	}

	return nil
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
	}

	return AllCommands
}
