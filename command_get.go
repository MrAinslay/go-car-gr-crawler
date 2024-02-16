package main

import (
	"log"

	"github.com/MrAinslay/go-car-gr-crawler/internal/client"
)

func commandGet(c client.Client, args ...string) error {
	url, err := client.GetUrl(args...)
	if err != nil {
		return err
	}

	log.Println(url)
	c.GetCarPosts(url)
	return nil
}
