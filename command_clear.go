package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MrAinslay/go-car-gr-crawler/internal/client"
)

func commandClear(c client.Client, args ...string) error {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Are you sure you want to DELETE ALL the results [Y]/[N]: ")
	for reader.Scan() {
		text := cleanInput(reader.Text())
		switch text {
		case "y", "yes":
			file, err := os.Open("results.json")
			if err != nil {
				return err
			}
			file.WriteString("{}")
			break
		case "n", "no":
			break
		default:
			fmt.Print("Invalid response: ")
		}
	}
	return nil
}
