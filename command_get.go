package main

import (
	"fmt"

	"github.com/MrAinslay/go-car-gr-crawler/internal/client"
)

func commandGet(c client.Client, args ...string) error {
	url, err := client.GetUrl(args...)
	if err != nil {
		return err
	}

	fmt.Printf("Visiting link: %s\n", url)
	return c.GetCarPosts(url)
}
