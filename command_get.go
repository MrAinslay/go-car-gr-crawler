package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MrAinslay/go-car-gr-crawler/internal/client"
)

func commandGet(c client.Client, args ...string) error {
	url := ""
	for _, arg := range args {
		if strings.Contains(arg, "limit=") {
			strLimit := strings.TrimPrefix(arg, "limit=")
			limit, err := strconv.Atoi(strLimit)
			if err != nil {
				fmt.Printf("Invalid limit defaulting to 1 err: %v", err)
				limit = 1
			}
			for i := range limit {
				url, err = client.GetUrl(fmt.Sprint(i), args...)
				if err != nil {
					return err
				}
				if err := c.GetCarPosts(url); err != nil {
					return err
				}
			}
			return nil
		}
	}

	url, err := client.GetUrl("", args...)
	if err != nil {
		return err
	}
	return c.GetCarPosts(url)
}
