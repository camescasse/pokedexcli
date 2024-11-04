package main

import (
	"time"

	"github.com/camescasse/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(time.Second*10, time.Minute*2)
	config := &config{
		pokeapiClient: client,
	}
	startRepl(config)
}
