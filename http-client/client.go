package http_client

import (
	"encoding/json"
	"errors"
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

	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		return getResponse, nil
	}

	return getResponse, errors.New("request failed")
}
