package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main entry point for the REPL tool, handles at high level the command and operations
func startRepl(cfg *config) {
	// scanner used to take in user input in the command line
	scanner := bufio.NewScanner(os.Stdin)

	// infinite for loop to keep the program running until killed or "exit" command
	for {
		fmt.Print("pokedex > ")

		// scan and convert input to string text
		scanner.Scan()
		text := scanner.Text()

		// convert input string to commands and arguments
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		// get command name from first index of cleaned
		commandName := cleaned[0]
		// remaining index of cleaned is args
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		// get and check if entered command is valid
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		// do actions of entered command and return error otherwise
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// struct to represent commands, callbacks of commands can be passed around as func
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	// a map of all available commands
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
		MAPBACK: {
			name:        MAPBACK,
			description: "Lists the previous page of location areas",
			callback:    callbackMapback,
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
	MAPBACK = "mapback"
	EXPLORE = "explore"
	CATCH   = "catch"
	INSPECT = "inspect"
	POKEDEX = "pokedex"
	EXIT    = "exit"
)
