package client

import (
	"fmt"

	"github.com/MrAinslay/go-car-gr-crawler/internal/parser"
)

func (c *Client) GetCarPosts(url string) error {
	if err := parser.Parse(url); err != nil {
		return err
	}
	fmt.Println("Saved results to results.json")
	return nil
}
