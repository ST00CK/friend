package service

import (
	"Friend/database"
	"Friend/dto"
	"Friend/repository"
	"context"
)

func CreateUserNode(user dto.UserDto) {
	neo4jUserRepository := repository.NewNeo4jUserRepository(database.Driver)

	_, err := neo4jUserRepository.CreateOrUpdateUserNode(context.Background(), user.UserID, user.Profile)
	if err != nil {
		return
	}
}

func DeleteUserNode(userID string) {
	neo4jUserRepository := repository.NewNeo4jUserRepository(database.Driver)

	_, err := neo4jUserRepository.DeleteUserNode(context.Background(), userID)
	if err != nil {
		return
	}
}

func CreateFriendship(user1ID string, user2ID string) {
	neo4jFriendRepository := repository.NewNeo4jUserRepository(database.Driver)

	_, err := neo4jFriendRepository.CreateUserRelation(context.Background(), user1ID, user2ID, "FRIEND")
	if err != nil {
		return
	}
}

func DeleteFriendship(user1ID string, user2ID string) {
	neo4jFriendRepository := repository.NewNeo4jUserRepository(database.Driver)

	_, err := neo4jFriendRepository.DeleteUserRelation(context.Background(), user1ID, user2ID, "FRIEND")
	if err != nil {
		return
	}
}

func GetFriends(userID string) (interface{}, error) {
	neo4jFriendRepository := repository.NewNeo4jUserRepository(database.Driver)

	res, err := neo4jFriendRepository.GetUserNode(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
