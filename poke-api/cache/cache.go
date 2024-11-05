package cache

import (
	"fmt"
	"maps"
	"time"
)

type cache struct {
	value     []byte
	createdAt time.Time
}

var memoryCache map[string]cache

func reapLoop(duration time.Duration) {
	fmt.Println("\n\ncleaning cache...")

	keys := maps.Keys(memoryCache)

	for key := range keys {
		cachedData := memoryCache[key]

		if time.Now().After(cachedData.createdAt.Add(duration)) {
			fmt.Printf("%v cache - removed\n", key)
			delete(memoryCache, key)
		}
	}

	fmt.Printf("\nPokedex CLI > ")
}

func Get(key string) ([]byte, bool) {
	cachedItem, ok := memoryCache[key]

	if ok {
		return cachedItem.value, true
	}

	return nil, false
}

func Add(key string, value []byte) {
	memoryCache[key] = cache{
		value:     value,
		createdAt: time.Now(),
	}
}

func NewCache(duration time.Duration) {
	memoryCache = make(map[string]cache)
	go func() {
		ticker := time.NewTicker(duration)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				reapLoop(duration)
			}
		}
	}()
}
