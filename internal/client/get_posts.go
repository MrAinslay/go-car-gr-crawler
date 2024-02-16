package client

import (
	"net/http"

	"golang.org/x/net/html"
)

func (c *Client) GetCarPosts(url string) error {
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
