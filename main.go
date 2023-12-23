package main

import (
	"log"
	"todo-api/controllers"
	"todo-api/initializers"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	router := server.Group("/tasks")
	router.GET("/", controllers.GetTasks)
	router.POST("/", controllers.CreateTask)
	router.GET("/:id", controllers.GetTaskById)

	log.Fatal(server.Run(":" + config.ServerPort))
}



