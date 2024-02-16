package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type VehicleListing struct {
	Name         string
	Price        string
	Link         string
	ImageURL     string
	Mileage      string
	CC           string
	Transmission string
	Fuel         string
	Date         string
}

func Parse(url string) error {
	writer, err := os.OpenFile("collector.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.car.gr"),
		colly.Debugger(&debug.LogDebugger{Output: writer}),
		colly.MaxDepth(2),
	)

	c.OnHTML("ol.list-unstyled.rows-container.mt-2.list.gallery-lg-4-per-row li", func(h *colly.HTMLElement) {
		vehicleDetails := formatText(h)
		fmt.Println(vehicleDetails)
	})

	if err := c.Visit(url); err != nil {
		return err
	}
	return nil
}

func formatText(h *colly.HTMLElement) []string {
	splitString := strings.Split(h.DOM.Text(), "\n")
	for i, line := range splitString {
		if strings.Contains(splitString[i], "χλμ") {
			features := strings.Split(splitString[i], ",")
			for n, feature := range features {
				features[n] = strings.TrimSpace(feature)
			}
			splitString[i] = ""

			splitString = append(splitString, features...)
		}
		splitString[i] = strings.TrimSpace(line)
	}
	return splitString
}
