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
		vehicle := formatTextToStruct(h)
		fmt.Printf("Name: %s\nLocation: %s\nDate: %s\nPrice: %s\nMileage: %s\nCC: %s\nHorsepower: %s\nFuel: %s\nLink: %s\n\n",
			vehicle.Name,
			vehicle.Location,
			vehicle.Date,
			vehicle.Price,
			vehicle.Mileage,
			vehicle.CC,
			vehicle.Horsepower,
			vehicle.Fuel,
			vehicle.Link,
		)
	})

	if err := c.Visit(url); err != nil {
		return err
	}
	return nil
}

func formatTextToStruct(h *colly.HTMLElement) VehicleListing {
	splitString := strings.Split(h.DOM.Text(), "\n")
	result := []string{}
	res := map[string]string{}
	for i, line := range splitString {
		if strings.Contains(splitString[i], "χλμ") {
			features := strings.Split(splitString[i], ",")
			for n, feature := range features {
				features[n] = strings.TrimSpace(feature)
			}
			splitString[i] = ""

			for _, feature := range features {
				switch {
				case strings.Contains(feature, "cc"):
					res["cc"] = feature
				case strings.Contains(feature, "bhp"):
					res["horsepower"] = feature
				case strings.Contains(feature, " χλμ"):
					res["mileage"] = feature
				case strings.Contains(feature, "/"):
					res["date"] = feature
				default:
					res["fuel"] = feature
				}
			}
		} else {
			splitString[i] = strings.TrimSpace(line)
		}
	}

	i := 0
	for _, line := range splitString {
		if line != "" && !strings.Contains(line, " / ") && !strings.Contains(line, "%") && !strings.Contains(line, "(Συζητήσιμη)") && !strings.Contains(line, "Με ζημιά") {
			switch i {
			case 0:
				res["name"] = line
			case 1:
				res["price"] = line
			case 2:
				res["location"] = line
			}
			i++
		}
	}
	res["link"] = fmt.Sprintf("https://www.car.gr%s", h.ChildAttr("a", "href"))

	return VehicleListing{
		Name: res["name"],
		Link: res["link"],
	}
}
