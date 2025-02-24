package controller

import (
	"Friend/dto"
	"Friend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserNode(c *gin.Context) {
	user := dto.UserDto{}
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "요청 형식이 잘못되었습니다."})
		return
	}

	service.CreateUserNode(user)

	c.JSON(http.StatusOK, gin.H{"message": "유저 노드 생성 완료"})
}

func DeleteUserNode(c *gin.Context) {
	user := dto.UserDto{}
	err := c.BindJSON(&user)

	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "요청 형식이 잘못되었습니다."})
		return
	}

	fmt.Println("userID: ", user.UserID)
	service.DeleteUserNode(user.UserID)

	c.JSON(http.StatusOK, gin.H{"message": "유저 노드 삭제 완료"})
}

func CreateFriendship(c *gin.Context) {
	friendship := dto.FriendshipDto{}
	err := c.BindJSON(&friendship)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "요청 형식이 잘못되었습니다."})
		return
	}

	service.CreateFriendship(friendship.User1ID, friendship.User2ID)

	c.JSON(http.StatusOK, gin.H{"message": "친구 관계 생성 완료"})
}

func DeleteFriendship(c *gin.Context) {
	friendship := dto.FriendshipDto{}
	err := c.BindJSON(&friendship)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "요청 형식이 잘못되었습니다."})
		return
	}

	service.DeleteFriendship(friendship.User1ID, friendship.User2ID)

	c.JSON(http.StatusOK, gin.H{"message": "친구 관계 삭제 완료"})
}

func GetFriends(c *gin.Context) {
	userID := c.Query("userID")
	res, err := service.GetFriends(userID)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "유저 노드를 찾을 수 없습니다."})
		return
	}

	c.JSON(http.StatusOK, res)
}
