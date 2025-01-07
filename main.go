package main

import (
	"Friend/config"
	"Friend/database"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	database.ConnectNeo4j()

	router := gin.Default()
	router.Run(":8080")
}
