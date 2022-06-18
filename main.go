package main

import (
	"log"

	"github.com/jardelkuhnen/video-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvironment()
	routes.Run()
}

func loadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
