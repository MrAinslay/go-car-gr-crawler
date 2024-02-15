package client

import (
	"net/http"

	htmlparser "github.com/MrAinslay/go-car-gr-crawler/internal/html_parser"
	"golang.org/x/net/html"
)

func (c *Client) GetCarPosts(searchQuery string, mileage string, sorting string) error {
	url, err := getUrl(searchQuery, mileage, sorting)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	htmlparser.ProcessAll(doc)
	return nil
}
