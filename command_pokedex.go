package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex:")
	// simply print out all entries that's stored in local variable
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
