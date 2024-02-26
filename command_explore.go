package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	// api call to get area info, list of available pokemons
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
