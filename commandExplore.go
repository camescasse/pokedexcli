package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("must provide a location area")
	}

	name := args[0]
	locationAreaDetails, err := config.pokeapiClient.GetLocationAreaDetails(name)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationAreaDetails.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
