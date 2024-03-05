package parser

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type VehicleListing struct {
	Name         string `json:"name" csv:"name"`
	Location     string `json:"location" csv:"location"`
	Link         string `json:"link" csv:"link"`
	Price        string `json:"price" csv:"price"`
	Date         string `json:"date" csv:"date"`
	Mileage      string `json:"mileage" csv:"mileage"`
	Horsepower   string `json:"horsepower" csv:"horsepower"`
	CC           string `json:"cc" csv:"cc"`
	Transmission string `json:"transmission" csv:"transmission"`
	Fuel         string `json:"fuel" csv:"fuel"`
}

type VehicleListingMap struct {
	Listings map[int]VehicleListing `json:"listing"`
}

type VehicleListingCSVMap struct {
	Listings [][]string
}

func Parse(url string) error {
	os.Mkdir("log", os.ModePerm)
	writer, err := os.OpenFile("log/collector.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.car.gr"),
		colly.Debugger(&debug.LogDebugger{Output: writer}),
		colly.MaxDepth(2),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting url: %s\n", url)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("err: %v\n", err)
	})

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

	listing := VehicleListing{
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

	switch {
	case listing.Name == "":
		listing.Name = "N/A"
	case listing.Link == "":
		listing.Link = "N/A"
	case listing.Location == "":
		listing.Location = "N/A"
	case listing.Date == "":
		listing.Date = "N/A"
	case listing.Horsepower == "":
		listing.Horsepower = "N/A"
	case listing.Mileage == "":
		listing.Mileage = "N/A"
	case listing.Price == "":
		listing.Price = "N/A"
	case listing.Fuel == "":
		listing.Fuel = "N/A"
	case listing.CC == "":
		listing.CC = "N/A"
	case listing.Transmission == "":
		listing.Transmission = "N/A"
	}

	return listing
}

func saveResults(vehicle VehicleListing) error {
	vehicleMap, err := saveJson(vehicle)
	if err != nil {
		return err
	}

	if err := saveCSV(vehicleMap); err != nil {
		return err
	}
	fmt.Println("Saved results to results.csv")
	return nil
}

func saveJson(vehicle VehicleListing) (VehicleListingMap, error) {
	writer, err := os.OpenFile("results.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return VehicleListingMap{}, err
	}

	dat, err := os.ReadFile("results.json")
	if err != nil {
		return VehicleListingMap{}, err
	}

	if string(dat) == "" {
		writer.WriteString("[]")
	}

	vehicles := VehicleListingMap{Listings: map[int]VehicleListing{}}
	if err := json.Unmarshal(dat, &vehicles.Listings); err != nil {
		return VehicleListingMap{}, err
	}

	for _, vch := range vehicles.Listings {
		if vch.Link == vehicle.Link {
			return VehicleListingMap{}, nil
		}
	}

	vehicles.Listings[len(vehicles.Listings)] = vehicle
	newDat, err := json.Marshal(vehicles.Listings)
	if err != nil {
		return VehicleListingMap{}, err
	}

	_, err = writer.Write(newDat)
	if err != nil {
		return VehicleListingMap{}, err
	}
	defer writer.Close()

	fmt.Println("Saved results to results.json")
	return vehicles, err
}

func saveCSV(vehicles VehicleListingMap) error {
	file, err := os.OpenFile("results.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	csvMap, err := listingsMapToCSV(vehicles)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	for _, record := range csvMap.Listings {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func listingsMapToCSV(m VehicleListingMap) (VehicleListingCSVMap, error) {
	result := VehicleListingCSVMap{}
	for _, vehicle := range m.Listings {
		newVehicleSlice := make([]string, 10)
		newVehicleSlice = append(newVehicleSlice,
			vehicle.Name,
			vehicle.Location,
			vehicle.Link,
			vehicle.Price,
			vehicle.Date,
			vehicle.Mileage,
			vehicle.Horsepower,
			vehicle.CC, vehicle.Transmission,
			vehicle.Fuel,
		)
		result.Listings = append(result.Listings, newVehicleSlice)
	}
	return result, nil
}
