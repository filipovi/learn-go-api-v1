package main

import (
	"log"

	"github.com/filipovi/learn-go-api-v1/app"
	"github.com/filipovi/learn-go-api-v1/config"
)

func main() {
	cfg, err := config.New("config.json")
	if nil != err {
		log.Fatal(err)
	}

	app.New(cfg)
}
