package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewReader(os.Stdin)
	commands := getCommands()

	for {
		fmt.Printf("pokedex > ")
		// os.Stdout.Sync()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			fmt.Println()
			continue
		}

		commandName := cleanInput(input)[0]
		if commandName == "" {
			continue
		}

		if command, exists := commands[commandName]; exists {
			err := command.callback()
			if err != nil {
				fmt.Fprintln(os.Stderr, "error executing command:", err)
				fmt.Println()
			}
		} else {
			fmt.Printf("unknown command: %s\ntype 'help' for available commands\n", input)
			fmt.Println()
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}

}
