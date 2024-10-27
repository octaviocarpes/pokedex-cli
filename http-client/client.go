package http_client

import (
	"encoding/json"
	"net/http"
)

const BASE_URL = "https://pokeapi.co/api/v2"

func ExecuteGet[T any](endpoint string) (T, error) {
	var getResponse T

	response, error := http.Get(BASE_URL + endpoint)

	if error != nil {
		return getResponse, error
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	decodeError := decoder.Decode(&getResponse)

	if decodeError != nil {
		return getResponse, decodeError
	}

	return getResponse, nil
}
