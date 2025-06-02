# Project Title

This project is a user management service with friendship functionality.

## Technologies Used

* Go
* Neo4j
* Docker

## Getting Started

### Prerequisites

* Docker
* Go

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up environment variables:
   Create a `.env` file in the root directory and add the following variables:
   ```
   NEO4J_URI=bolt://localhost:7687
   NEO4J_USERNAME=neo4j
   NEO4J_PASSWORD=password
   ```
4. Run the Neo4j database using Docker:
   ```bash
   docker run -d --name neo4j -p 7474:7474 -p 7687:7687 -e NEO4J_AUTH=neo4j/password neo4j:latest
   ```
5. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

### User

*   **POST /user/create**
    *   Description: Creates a new user node.
    *   Request body:
        ```json
        {
          "userID": "string",
          "name": "string"
        }
        ```
    *   Response:
        ```json
        {
          "message": "유저 노드 생성 완료"
        }
        ```
*   **POST /user/delete**
    *   Description: Deletes a user node.
    *   Request body:
        ```json
        {
          "userID": "string"
        }
        ```
    *   Response:
        ```json
        {
          "message": "유저 노드 삭제 완료"
        }
        ```
*   **POST /user/edit**
    *   Description: Edits a user node.
    *   Request body:
        ```json
        {
          "userID": "string",
          "name": "string"
        }
        ```
    *   Response:
        ```json
        {
          "message": "유저 노드 수정 완료"
        }
        ```

### Friendship

*   **POST /friendship/create**
    *   Description: Creates a friendship between two users.
    *   Request body:
        ```json
        {
          "user1ID": "string",
          "user2ID": "string"
        }
        ```
    *   Response:
        ```json
        {
          "message": "친구 관계 생성 완료"
        }
        ```
*   **POST /friendship/delete**
    *   Description: Deletes a friendship between two users.
    *   Request body:
        ```json
        {
          "user1ID": "string",
          "user2ID": "string"
        }
        ```
    *   Response:
        ```json
        {
          "message": "친구 관계 삭제 완료"
        }
        ```
*   **GET /friendship/list**
    *   Description: Lists all friends of a user.
    *   Query parameters:
        *   `userID`: string
    *   Response:
        ```json
        [
          {
            "userID": "string",
            "name": "string"
          }
        ]
        ```
