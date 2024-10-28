package main

import (
	"fmt"
	"os"

	"github.com/octaviocarpes/pokedex-cli/commands"
	"github.com/octaviocarpes/pokedex-cli/utils"
)

func main() {
	cliCommands := commands.ListAllCommands()
	var userInput string
	userInput = utils.RequestUserInput()

	for userInput != "exit" {

		command, ok := cliCommands[userInput]

		if ok {
			command.Callback()
			userInput = utils.RequestUserInput()
		} else {
			fmt.Println("Command not found - Exiting program")
			fmt.Println("*Tip type help to list all commands")
			os.Exit(1)
		}
	}
}
