package map_command

import (
	"fmt"
	"log"
	"os"

	poke_api "github.com/octaviocarpes/pokedex-cli/poke-api"
	utils "github.com/octaviocarpes/pokedex-cli/utils"
)

func printMapMode() {
	fmt.Printf("\n")
	fmt.Println("To load more maps type *next*")
	fmt.Println("To go back to previous maps type *back*")
	fmt.Println("To exit map search type *menu*")
	fmt.Println("To exit program type *exit cli*")
	fmt.Printf("\n\n")
}

func loadLocations(page int, preventLogs bool) int {
	response, error := poke_api.GetLocations(page)

	if error != nil {
		log.Fatal("Failed to fetch pokeapi locations")
		os.Exit(1)
	}

	if !preventLogs {
		fmt.Printf("Pokemon Locations: \n\n")

		for _, locataion := range response.Results {
			fmt.Printf("%s\n", locataion.Name)
		}
		fmt.Printf("\nBatch (%v)\n", page)
		printMapMode()
	}

	return response.Count
}

func navigateForward(page *int, offset int) {
	*page++
	loadLocations(*page*offset, false)
}

func GetLocations() error {
	const offset = 20
	var page int = 0

	loadLocations(0, false)

	userInput := utils.RequestUserInput()
	for {
		switch userInput {
		case "menu":
			utils.ClearConsole()
			return nil
		case "exit cli":
			os.Exit(0)
		case "next":
			navigateForward(&page, offset)
			userInput = utils.RequestUserInput()
		case "back":
			page--

			if page < 0 {
				page++
				fmt.Printf("\n")
				fmt.Println("Cannot go back before first page")
				userInput = utils.RequestUserInput()
				break
			}

			loadLocations(page*offset, false)
			userInput = utils.RequestUserInput()
		default:
			fmt.Printf("\n")
			fmt.Println("Unknown map mode command")
			printMapMode()
			userInput = utils.RequestUserInput()
		}
	}
}

func GetBackwardLocations() error {
	const offset = 20
	total := loadLocations(0, true)
	loadLocations(total-offset, false)

	userInput := utils.RequestUserInput()
	for {
		switch userInput {
		case "menu":
			utils.ClearConsole()
			return nil
		case "exit cli":
			os.Exit(0)
		case "next":
			total -= offset

			if total < 0 {
				total += offset
				fmt.Printf("\n")
				fmt.Println("Cannot go back before first page")
				userInput = utils.RequestUserInput()
				break
			}

			loadLocations(total, false)
			userInput = utils.RequestUserInput()

		case "back":
			total += offset
			navigateForward(&total, offset)
			userInput = utils.RequestUserInput()
		default:
			fmt.Printf("\n")
			fmt.Println("Unknown map mode command")
			printMapMode()
			userInput = utils.RequestUserInput()
		}
	}
}
