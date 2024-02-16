package client

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	baseURL = "https://www.car.gr/used-cars/"
)

func GetUrl(page string, args ...string) (url string, err error) {
	url = baseURL

	if page == "" {
		page = "1"
	}

	switch len(args) {
	case 0:
		err = errors.New("need atleast one argument")
	case 1:
		url = fmt.Sprintf("%s%s.html?activeq=%s&offer_type=sale&pg=%s&sort=pra", baseURL, args[0], args[0], page)
	case 2:
		mileage, err := strconv.Atoi(args[1])
		if err != nil {
			err = nil
			mileage = 150000
		}
		url = fmt.Sprintf("%s%s.html?activeq=%s&mileage-to=%d&offer_type=sale&pg=%s&sort=pra", baseURL, args[0], args[0], mileage, page)
	case 3:
		mileage, err := strconv.Atoi(args[1])
		if err != nil {
			err = nil
			mileage = 150000
		}
		url = fmt.Sprintf("%s%s.html?activeq=%s&mileage-to=%d&offer_type=sale&pg=%s&sort=%s", baseURL, args[0], args[0], mileage, page, args[2])
	case 4:
		mileage, err := strconv.Atoi(args[1])
		if err != nil {
			err = nil
			mileage = 150000
		}

		for _, arg := range args {
			if strings.Contains(arg, "limit=") {
				url = fmt.Sprintf("%s%s.html?activeq=%s&mileage-to=%d&offer_type=sale&pg=%s&sort=%s", baseURL, args[0], args[0], mileage, page, args[2])
			}
		}
		err = errors.New("too many arguments")
		_ = err
	default:
		err = errors.New("too many arguments")
	}
	return
}
