package main

import (
	"time"

	"github.com/camescasse/pokedexcli/internal/pokeapi"
)

func main() {

	const (
		requestTimeout = 10 * time.Second
		cacheInterval  = 2 * time.Minute
	)

	client := pokeapi.NewClient(requestTimeout, cacheInterval)
	config := &config{
		pokeapiClient: client,
		pokedex:       map[string]pokeapi.ResponsePokemon{},
	}
	startRepl(config)
}
