package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *config, args ...string) error {
	if len(config.pokedex) == 0 {
		return errors.New("no pokemon registered in pokedex")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
