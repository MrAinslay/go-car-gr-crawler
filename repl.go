package main

import (
	"strings"

	"github.com/MrAinslay/go-car-gr-crawler/internal/client"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c client.Client, s string) error
}

func cleanInput(s string) string {
	output := strings.TrimSpace(s)
	output = strings.ToLower(output)
	return output
}

func startRepl(c client.Client) {

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
		},
	}
}
