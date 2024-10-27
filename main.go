package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/octaviocarpes/pokedex-cli/commands"
)

func requestUserInput() string {
	fmt.Println("Pokedex CLI >")
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
		switch userInput {
		case "help":
			cliCommands["help"].Callback()
			userInput = requestUserInput()
		case "exit":
			fmt.Printf("exit: %v\n", cliCommands["exit"])
			userInput = requestUserInput()
		default:
			fmt.Printf("exit: %v\n", cliCommands["exit"])
			userInput = requestUserInput()
		}
	}
}
