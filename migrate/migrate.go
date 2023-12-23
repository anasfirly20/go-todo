package main

import (
	"fmt"
	"log"
	"todo-api/initializers"
	"todo-api/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Task{})
	fmt.Println("Migration complete")
}