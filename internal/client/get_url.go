package client

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	baseURL = "https://www.car.gr/used-cars/"
)

func GetUrl(args ...string) (url string, err error) {
	url = baseURL

	switch len(args) {
	case 0:
		err = errors.New("need atleast one argument")
	case 1:
		url = fmt.Sprintf("%s%s.html?activeq=%s&offer_type=sale&pg=1&sort=pra", baseURL, args[0], args[0])
	case 2:
		mileage, err := strconv.Atoi(args[1])
		if err != nil {
			err = nil
			mileage = 150000
		}
		url = fmt.Sprintf("%s%s.html?activeq=%s&mileage-to=%d&offer_type=sale&pg=1&sort=pra", baseURL, args[0], args[0], mileage)
	case 3:
		mileage, err := strconv.Atoi(args[1])
		if err != nil {
			err = nil
			mileage = 150000
		}
		url = fmt.Sprintf("%s%s.html?activeq=%s&mileage-to=%d&offer_type=sale&pg=1&sort=%s", baseURL, args[0], args[0], mileage, args[2])
	default:
		err = errors.New("too many arguments")
	}
	return
}
