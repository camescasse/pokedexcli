package main

import (
	"os"
)

func commandExit(config *config, parameter string) error {
	os.Exit(0)
	return nil
}
