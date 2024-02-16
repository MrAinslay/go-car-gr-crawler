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
	Mileage      string
	CC           string
	Horsepower   string
	Transmission string
	Fuel         string
	Date         string
	Location     string
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
		formatTextToStruct(h)
	})

	if err := c.Visit(url); err != nil {
		return err
	}
	return nil
}

func formatTextToStruct(h *colly.HTMLElement) VehicleListing {
	splitString := strings.Split(h.DOM.Text(), "\n")
	result := []string{}
	for i, line := range splitString {
		if strings.Contains(splitString[i], "χλμ") {
			features := strings.Split(splitString[i], ",")
			for n, feature := range features {
				features[n] = strings.TrimSpace(feature)
			}
			splitString[i] = ""

			splitString = append(splitString, features...)
		} else {
			splitString[i] = strings.TrimSpace(line)
		}
	}
	for _, line := range splitString {
		if line != "" && !strings.Contains(line, " / ") && !strings.Contains(line, "%") && !strings.Contains(line, "(Συζητήσιμη)") {
			result = append(result, line)
		}
	}
	result = append(result, fmt.Sprintf("https://www.car.gr%s", h.ChildAttr("a", "href")))

	return VehicleListing{
		Name:       result[0],
		Price:      result[1],
		Mileage:    result[2],
		CC:         result[3],
		Horsepower: result[4],
		Fuel:       result[5],
		Link:       result[len(result)-1],
	}
}
