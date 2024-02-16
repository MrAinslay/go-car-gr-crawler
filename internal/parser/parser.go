package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type VehicleListing struct {
	Name         string `json:"name"`
	Price        string `json:"price"`
	Link         string `json:"link"`
	Mileage      string `json:"mileage"`
	CC           string `json:"cc"`
	Horsepower   string `json:"horsepower"`
	Transmission string `json:"transmission"`
	Fuel         string `json:"fuel"`
	Date         string `json:"date"`
	Location     string `json:"location"`
}

type VehicleListingMap struct {
	Listings map[int]VehicleListing `json:"listing"`
}

func Parse(url string) error {
	os.Create("log")
	writer, err := os.OpenFile("log/collector.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.car.gr"),
		colly.Debugger(&debug.LogDebugger{Output: writer}),
		colly.MaxDepth(2),
	)

	c.OnHTML("ol.list-unstyled.rows-container.mt-2.list.gallery-lg-4-per-row li", func(h *colly.HTMLElement) {
		if err := saveResults(formatTextToStruct(h)); err != nil {
			log.Println(err)
		}
	})

	if err := c.Visit(url); err != nil {
		return err
	}
	return nil
}

func formatTextToStruct(h *colly.HTMLElement) VehicleListing {
	splitString := strings.Split(h.DOM.Text(), "\n")
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
				case strings.Contains(feature, " χλμ") && strings.Contains(feature, "."):
					res["mileage"] = feature
				case strings.Contains(feature, "/"):
					res["date"] = feature
				case strings.Contains(strings.ToLower(feature), "αυτόματο"):
					res["transmission"] = feature
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
		Name:         res["name"],
		Link:         res["link"],
		Location:     res["location"],
		Price:        res["price"],
		Date:         res["date"],
		Mileage:      res["mileage"],
		CC:           res["cc"],
		Horsepower:   res["horsepower"],
		Fuel:         res["fuel"],
		Transmission: res["transmission"],
	}
}

func saveResults(vehicle VehicleListing) error {
	writer, err := os.OpenFile("results.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	dat, err := os.ReadFile("results.json")
	if err != nil {
		return err
	}

	if string(dat) == "" {
		writer.WriteString("[]")
	}

	vehicles := VehicleListingMap{Listings: map[int]VehicleListing{}}
	if err := json.Unmarshal(dat, &vehicles.Listings); err != nil {
		return err
	}

	for _, vch := range vehicles.Listings {
		if vch.Link == vehicle.Link {
			return nil
		}
	}

	vehicles.Listings[len(vehicles.Listings)] = vehicle
	newDat, err := json.Marshal(vehicles.Listings)
	if err != nil {
		return err
	}

	_, err = writer.Write(newDat)
	defer writer.Close()
	return err
}
