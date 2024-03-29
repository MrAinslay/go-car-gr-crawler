package main

import (
	"fmt"

	"github.com/MrAinslay/go-car-gr-crawler/internal/client"
)

func commandHelp(c client.Client, args ...string) error {
	commands := getCommands()
	fmt.Println("A list of all commands:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("Command name: %s\nCommand Description: %s\n\n", cmd.name, cmd.description)
	}
	return nil
}
