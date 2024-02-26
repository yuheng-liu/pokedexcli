package main

import "os"

func callbackExit(cfg *config, args ...string) error {
	// simply exit when called
	os.Exit(0)
	return nil
}
