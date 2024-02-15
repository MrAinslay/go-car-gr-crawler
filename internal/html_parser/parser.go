package htmlparser

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func ProcessAll(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "li" {
		processNode(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ProcessAll(c)
	}
}

func processNode(n *html.Node) {
	switch n.Data {
	case "h2":
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			name := n.FirstChild.Data
			fmt.Println("\nName:", strings.TrimSpace(name))
		}
	case "span":
		for _, a := range n.Attr {
			switch a.Key {
			case "data-v-f1dc9bb4":
				price := ""
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "span" {
						price += c.FirstChild.Data
					}
				}
				if price != "" {
					fmt.Println("Price:", price)
				}

			case "title":
				fmt.Println("here")
				switch a.Val {
				case "Χρονολογία":
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.TextNode {
							fmt.Println("Date:", c.Data)
						}
					}
				}
			case "Χιλιόμετρα":
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						fmt.Println("Kilometers:", c.Data)
					}
				}
			case "Κυβικά":
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						fmt.Println("CC:", c.Data)
					}
				}
			case "Ίπποι":
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						fmt.Println("Horsepower:", c.Data)
					}
				}
			case "Σασμάν":
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						fmt.Println("Transmission:", c.Data)
					}
				}
			case "Καύσιμο":
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						fmt.Println("Fuel:", c.Data)
					}
				}
			}
		}
	case "img":
		for _, a := range n.Attr {
			if a.Key == "src" {
				ImageURL := a.Val
				fmt.Println("Image URL:", ImageURL)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processNode(c)
	}
}
