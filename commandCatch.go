package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	catchRate, err := calculateCatchRate(pokemon.BaseExperience)
	if err != nil {
		return err
	}

	odds := float64(rand.Intn(100) + 1)

	fmt.Println()
	fmt.Printf("Throwing Pokeball at %s...\n", pokemon.Name)
	if catchRate < odds {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		fmt.Println()
		return nil
	}

	config.pokedex[name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println()

	return nil
}
