package client

import (
	"github.com/MrAinslay/go-car-gr-crawler/internal/parser"
)

func (c *Client) GetCarPosts(url string) error {
	if err := parser.Parse(url); err != nil {
		return err
	}

	return nil
}
