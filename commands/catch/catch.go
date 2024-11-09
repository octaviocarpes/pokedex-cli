package catch_command

import (
	"fmt"
	"math/rand"
	"time"

	api "github.com/octaviocarpes/pokedex-cli/poke-api"
	pokedex "github.com/octaviocarpes/pokedex-cli/pokedex"
)

func weightedThrow(difficulty int) bool {
	baseProbability := 0.5

	// the higher the difficultyTuner, the easiest the throw gets
	difficultyTuner := 10
	probabilityOfThrow := baseProbability / (1 + float64(difficulty)/float64(difficultyTuner))

	// Set a minimum probability threshold (for example, 10%)
	if probabilityOfThrow < 0.1 {
		probabilityOfThrow = 0.1
	}

	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	if rand.Float64() < probabilityOfThrow {
		return false
	} else {
		return true
	}
}

func Catch(pokemon string) error {
	fmt.Printf("Throwing a pokeball at %v...\n", pokemon)

	response, err := api.GetPokemon(pokemon)

	if err != nil {
		return err
	}

	catchResult := weightedThrow(response.BaseExperience)

	if catchResult {
		fmt.Printf("Gotcha! %v was caught!\n", pokemon)
		pokedex.AddToPokedex(pokemon, response)
	} else {
		fmt.Printf("Oof! %v escaped the pokeball!\n", pokemon)
	}

	return nil
}
