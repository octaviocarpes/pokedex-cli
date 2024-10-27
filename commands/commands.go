package commands

import (
	"fmt"
	"os"
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
	`)

	return nil
}

func exitCallback() error {
	os.Exit(0)
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
	}

	return AllCommands
}
