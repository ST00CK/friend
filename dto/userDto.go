package dto

// UserDto UserDto는 외부 User DB로부터 받은 최소 정보(userID, profile)를 담는 구조체입니다.
type UserDto struct {
	UserID  string `json:"userID"`
	Profile string `json:"profile"`
}
