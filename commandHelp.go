package main

import (
	"fmt"
)

func commandHelp() error {
	commands := getCommands()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}
