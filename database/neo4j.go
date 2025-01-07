package database

import (
	"Friend/config"
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

var Driver neo4j.DriverWithContext

// ConnectNeo4j: Neo4j 드라이버 초기화 및 연결 확인

func ConnectNeo4j() {
	ctx := context.Background()

	uri := config.GetEnv("NEO4J_URI")
	username := config.GetEnv("NEO4J_USERNAME")
	password := config.GetEnv("NEO4J_PASSWORD")

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatalf("Neo4j 드라이버 생성 실패: %v", err)
	}

	if err = driver.VerifyConnectivity(ctx); err != nil {
		log.Fatalf("Neo4j 연결 실패: %v", err)
	}

	log.Println("Neo4j 연결 성공!")
	Driver = driver
}
