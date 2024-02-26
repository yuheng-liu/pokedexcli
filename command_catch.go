package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	// api call to simulate catching a pokemon
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// threshold to consider a pokemon is caught
	const threshold = 50
	// generate random number between 0 and a predefined value (BaseExperience) from api response
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	// only consider success if random number is lower than threshold
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	// add caught pokemon to variable list
	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
