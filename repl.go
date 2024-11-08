package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/camescasse/pokedexcli/internal/pokeapi"
)

func startRepl(config *config) {
	reader := bufio.NewReader(os.Stdin)
	commands := getCommands()

	for {
		fmt.Printf("pokedex > ")
		os.Stdout.Sync()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			fmt.Println()
			continue
		}

		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if commandName == "" {
			continue
		}

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if command, exists := commands[commandName]; exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println()
				fmt.Fprintln(os.Stderr, "error executing command:", err)
				fmt.Println()
			}
		} else {
			fmt.Println()
			fmt.Printf("unknown command: %s", input)
			fmt.Println("type 'help' for available commands")
			fmt.Println()
		}
	}
}

type config struct {
	pokedex       map[string]pokeapi.ResponsePokemon
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in Pokemon",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in Pokemon",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Explores a location area and shows the pokemon found in it",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "View details about a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your caught pokemon",
			callback:    commandPokedex,
		},
	}
}
