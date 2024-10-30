package cache

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type cache struct {
	value     []byte
	createdAt time.Time
}

var memoryCache map[string]cache = make(map[string]cache)

func Get(key string) any {
	cachedItem, ok := memoryCache[key]

	if ok {
		var cacheValue bytes.Buffer
		decoder := gob.NewDecoder(&cacheValue)
		error := decoder.Decode(&cachedItem.value)

		if error != nil {
			log.Fatal("Failed to decode cache value from byte slice")
		}

		return cacheValue
	}

	return nil
}

func Set(key string, value any) {
	var cacheValue bytes.Buffer
	encoder := gob.NewEncoder(&cacheValue)
	error := encoder.Encode(value)

	if error != nil {
		log.Fatal("Failed to encode cache value to byte slice")
	}

	memoryCache[key] = cache{
		value:     cacheValue.Bytes(),
		createdAt: time.Now(),
	}
}
