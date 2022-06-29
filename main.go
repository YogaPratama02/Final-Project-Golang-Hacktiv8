package main

import (
	"BootcampHacktiv8/final_project/db"
	"BootcampHacktiv8/final_project/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error read env file with err: %s", errEnv)
	}
	db.Connect()
	routes.Init()
}
