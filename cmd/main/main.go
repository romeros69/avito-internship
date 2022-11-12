package main

import (
	"avito-internship/configs"
	"log"
)

func main() {
	_, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}
}
