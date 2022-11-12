package main

import (
	"avito-internship/configs"
	"avito-internship/internal/myapp/app"
	"log"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}

	app.Run(cfg)
}
