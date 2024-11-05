package main

import (
	"errors"
	"fmt"
)

func commandMapf(config *config, parameter string) error {
	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.Next)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	return nil
}

func commandMapb(config *config, parameter string) error {
	if config.Previous == nil {
		return errors.New("no previous location areas to get")
	}

	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.Previous)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	return nil
}
