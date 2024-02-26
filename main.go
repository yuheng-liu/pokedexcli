package main

import (
	"time"

	"github.com/yuheng-liu/pokedexcli/internal/pokeapi"
)

// main struct that stores common info
type config struct {
	// basic client to make http requests
	pokeapiClient pokeapi.Client
	// used for pagination with map command
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	// used for storing pokemon info and to display pokedex
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	// initialization of the config struct
	// pointers can be null thus need not be initialized
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	// function to start the REPL(read-eval-print loop) in command line
	startRepl(&cfg)
}
