package client

import (
	"github.com/MrAinslay/go-car-gr-crawler/internal/parser"
)

func (c *Client) GetCarPosts(url string) error {
	return parser.Parse(url)
}
