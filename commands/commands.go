package commands

import (
	"fmt"
	"log"
	"os"

	poke_api "github.com/octaviocarpes/pokedex-cli/poke-api"
	utils "github.com/octaviocarpes/pokedex-cli/utils"
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

func printMapMode() {
	fmt.Printf("\n")
	fmt.Println("To load more maps type *next*")
	fmt.Println("To go back to previous maps type *back*")
	fmt.Println("To exit map search type *menu*")
	fmt.Println("To exit program type *exit cli*")
	fmt.Printf("\n\n")
}

func loadLocations(page int) int {
	response, error := poke_api.GetLocations(page)

	if error != nil {
		log.Fatal("Failed to fetch pokeapi locations")
		os.Exit(1)
	}

	fmt.Printf("Pokemon Locations: \n\n")

	for _, locataion := range response.Results {
		fmt.Printf("%s\n", locataion.Name)
	}
	fmt.Printf("\nBatch (%v)\n", page)
	printMapMode()

	return response.Count
}

// TODO: create map dedicated package
func mapCallback() error {
	const offset = 20
	var page int = 0

	loadLocations(0)

	userInput := utils.RequestUserInput()

	for {
		switch userInput {
		case "menu":
			fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			return nil
		case "exit cli":
			os.Exit(0)
		case "next":
			page++
			loadLocations(page * offset)
			userInput = utils.RequestUserInput()
		case "back":
			page--

			if page <= 0 {
				page++
				fmt.Printf("\n")
				fmt.Println("Cannot go back before first page")
				userInput = utils.RequestUserInput()
				break
			}

			loadLocations(page * offset)
			userInput = utils.RequestUserInput()
		default:
			fmt.Printf("\n")
			fmt.Println("Unknown map mode command")
			printMapMode()
			userInput = utils.RequestUserInput()
		}
	}
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
