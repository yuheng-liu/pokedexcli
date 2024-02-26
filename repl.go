package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableComands := getCommands()

		command, ok := availableComands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		HELP: {
			name:        HELP,
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		MAP: {
			name:        MAP,
			description: "Lists the next page of location areas",
			callback:    callbackMap,
		},
		MAPB: {
			name:        MAPB,
			description: "Lists the previous page of location areas",
			callback:    callbackMapb,
		},
		EXPLORE: {
			name:        EXPLORE + " {location_name}",
			description: "Lists the pokemon in a location area",
			callback:    callbackExplore,
		},
		CATCH: {
			name:        CATCH + " {pokemon_name}",
			description: "Attempty to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		INSPECT: {
			name:        INSPECT + " {pokemon_name}",
			description: "View information about a caught pokemon",
			callback:    callbackInpect,
		},
		POKEDEX: {
			name:        POKEDEX,
			description: "View all the pokemon in your pokedex",
			callback:    callbackPokedex,
		},
		EXIT: {
			name:        EXIT,
			description: "Turns off the Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

const (
	HELP    = "help"
	MAP     = "map"
	MAPB    = "mapb"
	EXPLORE = "explore"
	CATCH   = "catch"
	INSPECT = "inspect"
	POKEDEX = "pokedex"
	EXIT    = "exit"
)
