# 프로젝트 제목

이 프로젝트는 사용자 관리 및 친구 관계 기능을 제공하는 서비스입니다. 사용자를 생성, 수정, 삭제하고 사용자 간의 친구 관계를 맺고 끊을 수 있는 RESTful API를 제공합니다.

## 사용 기술

*   **Go:** 백엔드 API 개발에 사용된 프로그래밍 언어입니다.
*   **Neo4j:** 사용자 및 친구 관계 데이터를 저장하는 데 사용되는 그래프 데이터베이스입니다.
*   **Docker:** 애플리케이션 및 데이터베이스를 컨테이너화하여 실행 환경을 일관되게 유지하는 데 사용됩니다.
*   **Gin Gonic:** Go 언어 기반의 HTTP 웹 프레임워크로, API 라우팅 및 요청 처리를 담당합니다.

## 시작하기

### 사전 준비 사항

*   **Docker:** Neo4j 데이터베이스를 실행하기 위해 필요합니다. [Docker 설치 가이드](https://docs.docker.com/get-docker/)
*   **Go:** 애플리케이션을 실행하기 위해 필요합니다. (버전 1.18 이상 권장) [Go 설치 가이드](https://golang.org/doc/install)

### 설치 및 실행

1.  **프로젝트 저장소 복제:**
    ```bash
    git clone <프로젝트_저장소_URL>
    cd <프로젝트_디렉토리>
    ```
2.  **Go 모듈 종속성 설치:**
    프로젝트에 필요한 Go 라이브러리들을 다운로드하고 설치합니다.
    ```bash
    go mod download
    ```
3.  **환경 변수 설정:**
    프로젝트 루트 디렉토리에 `.env` 파일을 생성하고 다음과 같이 Neo4j 데이터베이스 연결 정보를 입력합니다. 이미 파일이 있다면 해당 내용을 업데이트합니다.
    ```
    NEO4J_URI=bolt://localhost:7687
    NEO4J_USERNAME=neo4j
    NEO4J_PASSWORD=your_neo4j_password
    ```
    *   `NEO4J_URI`: Neo4j 데이터베이스 서버 주소입니다. Docker를 로컬에서 실행하는 경우 기본값은 `bolt://localhost:7687` 입니다.
    *   `NEO4J_USERNAME`: Neo4j 데이터베이스 사용자 이름입니다.
    *   `NEO4J_PASSWORD`: Neo4j 데이터베이스 비밀번호입니다. Docker 실행 시 설정한 비밀번호와 일치해야 합니다.

4.  **Neo4j 데이터베이스 실행 (Docker 사용):**
    Docker를 사용하여 Neo4j 데이터베이스 인스턴스를 실행합니다. 이미 실행 중인 `neo4j` 컨테이너가 있다면 중지 후 삭제하고 다시 실행하거나, 다른 이름으로 실행하십시오.
    ```bash
    docker run -d --name neo4j-friend-service -p 7474:7474 -p 7687:7687 -e NEO4J_AUTH=neo4j/your_neo4j_password neo4j:latest
    ```
    *   `-d`: 백그라운드에서 컨테이너를 실행합니다.
    *   `--name neo4j-friend-service`: 컨테이너의 이름을 지정합니다.
    *   `-p 7474:7474`: Neo4j 브라우저 인터페이스를 위한 포트 매핑입니다. (호스트 포트:컨테이너 포트)
    *   `-p 7687:7687`: Bolt 프로토콜을 위한 포트 매핑입니다. (애플리케이션 연결용)
    *   `-e NEO4J_AUTH=neo4j/your_neo4j_password`: Neo4j 데이터베이스의 사용자 이름과 비밀번호를 설정합니다. 여기서 `your_neo4j_password` 부분을 `.env` 파일에 설정한 `NEO4J_PASSWORD`와 동일하게 변경해야 합니다.
    *   `neo4j:latest`: 최신 버전의 Neo4j 이미지를 사용합니다.

    Neo4j 브라우저는 `http://localhost:7474` 에서 접근할 수 있습니다.

5.  **애플리케이션 실행:**
    애플리케이션 서버를 시작합니다.
    ```bash
    go run main.go
    ```
    애플리케이션이 성공적으로 실행되면, API 요청을 받을 준비가 된 것입니다. 기본적으로 `http://localhost:8080` (설정에 따라 다를 수 있음)에서 실행됩니다.

## API 엔드포인트

API는 JSON 형식을 사용하여 요청 및 응답을 처리합니다.

### 사용자 (User)

*   **POST /user/create**
    *   설명: 새로운 사용자 노드를 생성합니다.
    *   요청 본문 (Request Body):
        ```json
        {
          "userID": "string (필수)",
          "name": "string (필수)"
        }
        ```
        *   `userID`: 사용자의 고유 아이디입니다.
        *   `name`: 사용자의 이름입니다.
    *   성공 응답 (Success Response) - 상태 코드 200:
        ```json
        {
          "message": "유저 노드 생성 완료"
        }
        ```
    *   실패 응답 (Error Response) - 상태 코드 400 (잘못된 요청):
        ```json
        {
          "error": "요청 형식이 잘못되었습니다."
        }
        ```
*   **POST /user/delete**
    *   설명: 기존 사용자 노드를 삭제합니다. 해당 사용자와 연결된 모든 친구 관계도 함께 삭제됩니다.
    *   요청 본문 (Request Body):
        ```json
        {
          "userID": "string (필수)"
        }
        ```
        *   `userID`: 삭제할 사용자의 고유 아이디입니다.
    *   성공 응답 (Success Response) - 상태 코드 200:
        ```json
        {
          "message": "유저 노드 삭제 완료"
        }
        ```
    *   실패 응답 (Error Response) - 상태 코드 400 (잘못된 요청):
        ```json
        {
          "error": "요청 형식이 잘못되었습니다."
        }
        ```
*   **POST /user/edit**
    *   설명: 기존 사용자 노드의 정보를 수정합니다. `userID`는 변경할 수 없으며, `name`만 변경 가능합니다.
    *   요청 본문 (Request Body):
        ```json
        {
          "userID": "string (필수)",
          "name": "string (필수, 변경할 이름)"
        }
        ```
        *   `userID`: 정보를 수정할 사용자의 고유 아이디입니다.
        *   `name`: 변경할 사용자의 새 이름입니다.
    *   성공 응답 (Success Response) - 상태 코드 200:
        ```json
        {
          "message": "유저 노드 수정 완료"
        }
        ```
    *   실패 응답 (Error Response) - 상태 코드 400 (잘못된 요청):
        ```json
        {
          "error": "요청 형식이 잘못되었습니다."
        }
        ```

### 친구 관계 (Friendship)

*   **POST /friendship/create**
    *   설명: 두 사용자 간에 친구 관계를 생성합니다. 이미 친구 관계이거나 존재하지 않는 사용자인 경우 오류가 발생할 수 있습니다.
    *   요청 본문 (Request Body):
        ```json
        {
          "user1ID": "string (필수)",
          "user2ID": "string (필수)"
        }
        ```
        *   `user1ID`: 친구 관계를 맺을 첫 번째 사용자의 아이디입니다.
        *   `user2ID`: 친구 관계를 맺을 두 번째 사용자의 아이디입니다. (user1ID와 달라야 합니다)
    *   성공 응답 (Success Response) - 상태 코드 200:
        ```json
        {
          "message": "친구 관계 생성 완료"
        }
        ```
    *   실패 응답 (Error Response) - 상태 코드 400 (잘못된 요청):
        ```json
        {
          "error": "요청 형식이 잘못되었습니다."
        }
        ```
        (또는 서비스 로직에 따른 다른 오류 메시지, 예를 들어 "이미 친구 관계입니다." 또는 "사용자를 찾을 수 없습니다.")
*   **POST /friendship/delete**
    *   설명: 두 사용자 간의 친구 관계를 삭제합니다.
    *   요청 본문 (Request Body):
        ```json
        {
          "user1ID": "string (필수)",
          "user2ID": "string (필수)"
        }
        ```
        *   `user1ID`: 친구 관계를 끊을 첫 번째 사용자의 아이디입니다.
        *   `user2ID`: 친구 관계를 끊을 두 번째 사용자의 아이디입니다.
    *   성공 응답 (Success Response) - 상태 코드 200:
        ```json
        {
          "message": "친구 관계 삭제 완료"
        }
        ```
    *   실패 응답 (Error Response) - 상태 코드 400 (잘못된 요청):
        ```json
        {
          "error": "요청 형식이 잘못되었습니다."
        }
        ```
*   **GET /friendship/list**
    *   설명: 특정 사용자의 모든 친구 목록을 조회합니다.
    *   쿼리 파라미터 (Query Parameters):
        *   `userID`: string (필수) - 친구 목록을 조회할 사용자의 아이디입니다.
        *   예시: `/friendship/list?userID=userA`
    *   성공 응답 (Success Response) - 상태 코드 200:
        ```json
        [
          {
            "userID": "string",
            "name": "string"
          },
          {
            "userID": "string",
            "name": "string"
          }
        ]
        ```
        (친구가 없는 경우 빈 배열 `[]`을 반환합니다.)
    *   실패 응답 (Error Response) - 상태 코드 400 (잘못된 요청 또는 사용자를 찾을 수 없음):
        ```json
        {
          "error": "유저 노드를 찾을 수 없습니다."
        }
        ```
