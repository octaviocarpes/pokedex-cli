package main

import (
	"fmt"
	"os"
	"time"

	"github.com/octaviocarpes/pokedex-cli/commands"
	"github.com/octaviocarpes/pokedex-cli/poke-api/cache"
	"github.com/octaviocarpes/pokedex-cli/utils"
)

func main() {
	cache.NewCache(30 * time.Second)
	cliCommands := commands.ListAllCommands()
	userInput, args := utils.RequestUserInput()

	for userInput != "exit" {

		command, ok := cliCommands[userInput]

		if ok {
			command.Callback(args)
			userInput, args = utils.RequestUserInput()
		} else {
			fmt.Printf("Command not found (%v) - Exiting program\n", userInput)
			fmt.Println("*Tip type help to list all commands")
			os.Exit(1)
		}
	}
}
