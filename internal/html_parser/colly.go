package htmlparser

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func Scrape(url string) {
	c := colly.NewCollector()

	c.Visit(url)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Printf("%v\n", e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	var posts []Post
	c.OnHTML("li", func(h *colly.HTMLElement) {
		carPost := Post{}

		carPost.Name = h.ChildText("h2")
		carPost.URL = h.ChildAttr("a", "href")
		h.ForEachWithBreak("span", func(i int, n *colly.HTMLElement) bool {
			if n.Attr("data-v-f1dc9bb4") != "" {
				carPost.Price += n.ChildText("span")
				return false
			}
			if n.Attr("") == "" {

			}
			return true
		})
	})

}
