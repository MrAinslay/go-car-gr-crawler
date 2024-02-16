package main

import (
	"os"

	"github.com/MrAinslay/go-car-gr-crawler/internal/client"
)

func commandExit(c client.Client, args ...string) error {
	os.Exit(1)
	return nil
}
