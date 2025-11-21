package main

import (
	"log"

	"task-manager-fullstack/config"
	"task-manager-fullstack/routes"

	"github.com/joho/godotenv"
)

func main() {
	var err = godotenv.Load()
	if err != nil {
		log.Fatal("env not found")
	}
	config.DbConnect()
	log.Println("DB after connect:", config.DB)

	r := routes.SetupRoutes()

	r.Run()
}
