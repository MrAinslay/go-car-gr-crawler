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
outer:
	for reader.Scan() {
		text := cleanInput(reader.Text())
		switch text {
		case "y", "yes":
			os.Remove("results.json")
			file, err := os.OpenFile("results.json", os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				return err
			}
			defer file.Close()

			file.Write([]byte("{}"))
			fmt.Println("Successfully cleared results file")
			break outer
		case "n", "no":
			break outer
		default:
			fmt.Print("Invalid response: ")
		}
	}
	return nil
}
