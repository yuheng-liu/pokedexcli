package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	// retrieve locations via api call
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	// used for pagination of locations results
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func callbackMapback(cfg *config, args ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	// retrieve locations via api call
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	// used for pagination of locations results
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}
