package repository

import "context"

// UserRepository UserRepository는 User 노드 생성/업데이트 등의 동작을 추상화합니다.
type UserRepository interface {
	// CreateOrUpdateUserNode 외부 User DB로부터 받은 최소 정보(userID, profile)를 기반으로 User 노드를 생성하거나 업데이트합니다.
	CreateOrUpdateUserNode(ctx context.Context, userID string, profile string) (interface{}, error)

	// DeleteUserNode User 노드를 삭제합니다.
	DeleteUserNode(ctx context.Context, userID string) (interface{}, error)

	// GetUserNode User 노드를 조회합니다.
	GetUserNode(ctx context.Context, userID string) (interface{}, error)

	// CreateUserRelation User 노드의 관계를 생성합니다.
	CreateUserRelation(ctx context.Context, userID string, targetUserID string, relation string) (interface{}, error)

	// DeleteUserRelation User 노드의 관계를 삭제합니다.
	DeleteUserRelation(ctx context.Context, userID string, targetUserID string, relation string) (interface{}, error)
}
