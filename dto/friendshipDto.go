package dto

// FriendshipDto is a struct that represents the friendship relationship between two users.
type FriendshipDto struct {
	User1ID string `json:"user1Id"`
	User2ID string `json:"user2Id"`
}
