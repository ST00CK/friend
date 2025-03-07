package main

import (
	"Friend/config"
	"Friend/database"
	"Friend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectNeo4j()
	router := gin.Default()
	router.Use(cors.New(config.SetCors()))
	routes.SetupRoutes(router)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
