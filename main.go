package main

import (
	"log"
	"todo-api/controllers"
	"todo-api/initializers"
	"todo-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	TaskController controllers.TaskController
	TaskRouteController routes.TaskRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	TaskController = controllers.NewTaskController(initializers.DB)
	TaskRouteController = routes.NewTaskRouteController(TaskController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
	
	router := server.Group("/api")
	
	TaskRouteController.TaskRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}