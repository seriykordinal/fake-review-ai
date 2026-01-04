package main

import (
	"log"

	"github.com/seriykordinal/fake-review-ai/internal/app"
	"github.com/seriykordinal/fake-review-ai/internal/config"
)

func main() {
	log.Println("--------fake-review-ai--------")

	cfg := config.Load()

	server, err := app.NewServer(cfg)

	if err != nil {
		log.Fatal(err)
	}

	server.Run()
}
