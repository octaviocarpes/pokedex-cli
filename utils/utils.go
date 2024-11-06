package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func RequestUserInput() (string, []string) {
	fmt.Printf("\nPokedex CLI > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	args := strings.Fields(scanner.Text())

	return args[0], args[1:]
}

func ClearConsole() {
	for i := 0; i < 50; i++ {
		fmt.Printf("\n")
	}
}

func ConvertToByteSlice(payload any) []byte {
	result, marshalError := json.Marshal(payload)

	if marshalError != nil {
		return nil
	}

	return result
}

func ConvertFromByteSlice[T any](payload []byte) (T, error) {
	var result T
	unmarshalError := json.Unmarshal(payload, &result)

	if unmarshalError != nil {
		return result, unmarshalError
	}

	return result, nil
}
