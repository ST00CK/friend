package repository

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type neo4jUserRepository struct {
	driver neo4j.DriverWithContext
}

// NewNeo4jUserRepository NewNeo4jUserRepository는 Neo4j 드라이버를 주입받아 UserRepository 인터페이스를 반환합니다.
func NewNeo4jUserRepository(driver neo4j.DriverWithContext) UserRepository {
	return &neo4jUserRepository{driver: driver}
}

// CreateOrUpdateUserNode CreateOrUpdateUserNode는 외부 User DB에서 전달받은 userID와 profile 데이터를
// 기반으로 User 노드를 생성(MERGE 사용)하거나 업데이트합니다.
func (r *neo4jUserRepository) CreateOrUpdateUserNode(ctx context.Context, userID string, profile string) (interface{}, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		// userID 기준으로 노드가 없으면 생성, 존재하면 profile 업데이트
		query := `
			MERGE (u:User {id: $userID})
			ON CREATE SET u.profile = $profile
			ON MATCH SET u.profile = $profile
			RETURN u
		`
		params := map[string]interface{}{
			"userID":  userID,
			"profile": profile,
		}

		res, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		record, err := res.Single(ctx)
		if err != nil {
			return nil, err
		}

		// record.Values[0]에 생성(또는 조회)된 노드가 담겨 있음
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *neo4jUserRepository) DeleteUserNode(ctx context.Context, userID string) (interface{}, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		// userID 기준으로 노드가 없으면 생성, 존재하면 profile 업데이트
		query := `
			MATCH (u:User {id: $userID})
			DETACH DELETE u
		`
		params := map[string]interface{}{
			"userID": userID,
		}

		res, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		record, err := res.Single(ctx)
		if err != nil {
			return nil, err
		}

		// record.Values[0]에 생성(또는 조회)된 노드가 담겨 있음
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *neo4jUserRepository) GetUserNode(ctx context.Context, userID string) (interface{}, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		// userID 기준으로 노드가 없으면 생성, 존재하면 profile 업데이트
		query := `
			MATCH (me:User {id: $userID})-[:FRIEND]->(friend:User)
			RETURN friend
		`
		params := map[string]interface{}{
			"userID": userID,
		}

		res, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		record, err := res.Single(ctx)
		if err != nil {
			return nil, err
		}

		// record.Values[0]에 생성(또는 조회)된 노드가 담겨 있음
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *neo4jUserRepository) CreateUserRelation(ctx context.Context, userID string, targetUserID string, relation string) (interface{}, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		// userID 기준으로 노드가 없으면 생성, 존재하면 profile 업데이트
		query := `
			MATCH (u:User {id: $userID})
			MATCH (t:User {id: $targetUserID})
			MERGE (u)-[r: FRIEND]->(t)
			RETURN r
		`
		params := map[string]interface{}{
			"userID":       userID,
			"targetUserID": targetUserID,
		}

		res, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		record, err := res.Single(ctx)
		if err != nil {
			return nil, err
		}

		// record.Values[0]에 생성(또는 조회)된 노드가 담겨 있음
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *neo4jUserRepository) DeleteUserRelation(ctx context.Context, userID string, targetUserID string, relation string) (interface{}, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		// userID 기준으로 노드가 없으면 생성, 존재하면 profile 업데이트
		query := `
			MATCH (u:User {id: $userID})
			MATCH (t:User {id: $targetUserID})
			MATCH (u)-[r:FRIEND]->(t)
			DELETE r
		`
		params := map[string]interface{}{
			"userID":       userID,
			"targetUserID": targetUserID,
			"relation":     relation,
		}

		res, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		record, err := res.Single(ctx)
		if err != nil {
			return nil, err
		}

		// record.Values[0]에 생성(또는 조회)된 노드가 담겨 있음
		return record.Values[0], nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
