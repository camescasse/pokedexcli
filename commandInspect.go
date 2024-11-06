package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := config.pokedex[name]
	if !ok {
		return errors.New("pokemon not yet caught")
	}

	fmt.Println()
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}
	fmt.Println("Stats:")
	for _, pokemonStat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", pokemonStat.Stat.Name, pokemonStat.BaseStat)
	}
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println()

	return nil
}
