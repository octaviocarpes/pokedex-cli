package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/octaviocarpes/pokedex-cli/commands"
)

func requestUserInput() string {
	fmt.Printf("\nPokedex CLI > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return scanner.Text()
}

func main() {
	cliCommands := commands.ListAllCommands()
	var userInput string
	userInput = requestUserInput()

	for userInput != "exit" {

		command, ok := cliCommands[userInput]

		if ok {
			command.Callback()
			userInput = requestUserInput()
		} else {
			fmt.Println("Command not found - Exiting program")
			fmt.Println("*Tip type help to list all commands")
			os.Exit(1)
		}
	}
}
