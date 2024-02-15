package main

import (
	"flag"
	"time"

	"github.com/MrAinslay/go-car-gr-crawler/internal/client"
)

func main() {
	timeout := flag.Int("t", 5, "client timeout")
	flag.Parse()

	client := client.New(time.Duration(*timeout) * time.Second)
	client.GetCarPosts("audi", "200000", "pra")
	startRepl(client)
}
