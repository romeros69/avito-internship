package main

import (
	"avito-internship/configs"
	"avito-internship/internal/myapp/app"
	"log"
)

// @tittle Avito test task by Roman Loginov
// @version 1.0
// @description Avito test task by Roman Loginov

// @host localhost:9000
// @BasePath /
func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}

	app.Run(cfg)
}
