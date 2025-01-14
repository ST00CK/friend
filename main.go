package main

import (
	"Friend/database"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectNeo4j()

	router := gin.Default()

	router.Run(":8080")

}
