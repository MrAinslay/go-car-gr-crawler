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
	callback    func(c client.Client, args ...string) error
}

func cleanInput(s string) string {
	output := strings.TrimSpace(s)
	output = strings.ToLower(output)
	return output
}

func printPrompt() {
	fmt.Print("Car-gr Scraper > ")
}

func startRepl(c client.Client) {
	commands := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		splitText := strings.Split(text, " ")
		if command, exists := commands[splitText[0]]; exists {
			err := command.callback(c, splitText[1:]...)
			if err != nil {
				log.Printf("err: %v", err)
			}
		} else {
			fmt.Println("Unknown command\nUse command help for help")
		}
		printPrompt()
	}
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
			description: "Exits the CLI",
			callback:    commandExit,
		},
		"get": {
			name:        "get",
			description: "Saves all the listings found in pages\nUsage: get <CAR_NAME> <MILEAGE> <ORDER_BY> <limit=<LIMIT>> - mileage, limit and order by are optional",
			callback:    commandGet,
		},
		"clear": {
			name:        "clear",
			description: "Clears the results file",
			callback:    commandClear,
		},
	}
}
