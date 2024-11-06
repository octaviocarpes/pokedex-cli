package cache

import (
	"log"

	utils "github.com/octaviocarpes/pokedex-cli/utils"
)

func GetCachedData[T any](url string) (T, error) {
	cachedResponse, ok := Get(url)
	var response T

	if ok {
		parsedResponse, conversionError := utils.ConvertFromByteSlice[T](cachedResponse)

		if conversionError != nil {
			log.Fatal("Failed to decode data from byte slice")
		}

		return parsedResponse, nil
	}

	return response, nil
}

func SaveDataInCache[T any](url string, payload T) bool {
	byteSlice := utils.ConvertToByteSlice(payload)

	if byteSlice != nil {
		Add(url, byteSlice)
	}

	return true
}
