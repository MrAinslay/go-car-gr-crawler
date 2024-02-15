package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MrAinslay/go-car-gr-crawler/internal/client"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c client.Client, args []string) error
}

func cleanInput(s string) string {
	output := strings.TrimSpace(s)
	output = strings.ToLower(output)
	return output
}

func printPrompt() {
	fmt.Println("Car-gr Scraper > ")
}

func startRepl(c client.Client) {
	commands := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		splitText := strings.Split(text, " ")
		if command, exists := commands[splitText[0]]; exists {
			err := command.callback(c, splitText[1:])
			if err != nil {
				log.Printf("err: %v", err)
			}
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
