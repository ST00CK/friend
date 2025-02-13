package routes

import (
	"Friend/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/user/create", controller.CreateUserNode)
	router.POST("/friendship/create", controller.CreateFriendship)
	router.POST("/friendship/delete", controller.DeleteFriendship)
	router.GET("/friendship/list", controller.GetFriends)
}
