package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RequestUserInput() string {
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
