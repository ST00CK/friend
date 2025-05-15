# **시스템 아키텍처 개요**
# **소개**

Friend 애플리케이션은 계층화된 아키텍처 패턴을 따르며, 다양한 계층에서 관심사를 분리합니다:

    API 계층: Gin 웹 프레임워크를 사용하여 HTTP 요청 및 응답 처리
    비즈니스 계층: 서비스 컴포넌트를 통한 애플리케이션 로직 구현
    데이터 접근 계층: 저장소를 통한 데이터베이스 작업 관리

이러한 분리는 코드베이스의 유지보수성, 테스트 가능성 및 유연성을 향상시킵니다.
# **아키텍처 계층** 

![image](https://github.com/user-attachments/assets/5d7ee54e-496d-440d-abc5-f9a2056b515f)

# **컴포넌트 관계**

다음 다이어그램은 시스템의 주요 컴포넌트와 파일 간의 관계를 보여줍니다:

![image](https://github.com/user-attachments/assets/1ab9b5a3-22a5-42cc-9669-889ecf44907c)

# **요청 처리 흐름** 

다음 시퀀스 다이어그램은 일반적인 요청이 시스템을 통해 어떻게 흐르는지 보여줍니다:

![image](https://github.com/user-attachments/assets/12260561-263e-4e56-87cd-d33a06be28fa)

# **핵심 컴포넌트**

## **메인 애플리케이션**

애플리케이션 진입점(main.go)은 다음과 같은 중요한 초기화 작업을 수행합니다:

    환경 설정 로드
    Neo4j 데이터베이스 연결 설정
    CORS 지원을 통한 Gin 웹 프레임워크 설정
    API 라우트 구성
    8080 포트에서 HTTP 서버 시작 

## **컨트롤러**

컨트롤러 계층은 HTTP 요청을 처리하고 이를 서비스 호출로 변환합니다. 주요 컨트롤러는 userController.go로, 다음 엔드포인트를 관리합니다:
| 엔드포인트             | 메소드 | 설명                     | 코드 참조 
|------------------------|--------|--------------------------|--------------------------------|
| `/user/create`         | POST   | 새 사용자 노드 생성      | `userController.go:11-23`      |                         
| `/user/delete`         | POST   | 사용자 노드 삭제         | `userController.go:25-39`      |                         
| `/user/edit`           | POST   | 사용자 프로필 업데이트   | `userController.go:41-53`      |                         
| `/friendship/create`   | POST   | 사용자 간 친구 관계 생성 | `userController.go:55-67`      |                         
| `/friendship/delete`   | POST   | 친구 관계 제거           | `userController.go:69-81`      |  
| `/friends`             | GET    | 사용자의 친구 목록 조회  | `userController.go:83-93`      |

## **서비스**

서비스 계층은 컨트롤러와 저장소 간의 작업을 조정하는 비즈니스 로직을 포함합니다:

![image](https://github.com/user-attachments/assets/4c76b4e8-a4ba-4bf9-9018-5a5a8575c1d5)

## **저장소**

저장소 계층은 Neo4j Go 드라이버를 통해 Neo4j 데이터베이스에 대한 데이터 접근을 처리합니다:

    Neo4jUserRepository: UserRepository 인터페이스 구현
    Neo4j에 대해 CRUD 작업을 수행하기 위한 Cypher 쿼리 실행
    데이터베이스 상호작용을 위한 트랜잭션 기반 세션 관리

주요 작업:

    사용자 노드 관리 (생성, 업데이트, 삭제)
    관계 관리 (친구 연결 생성, 삭제)
    사용자 관계 쿼리 (친구 찾기

# **데이터 모델**

애플리케이션은 Neo4j에서 다음과 같은 구조의 그래프 데이터 모델을 사용합니다:

![image](https://github.com/user-attachments/assets/62499f26-4b6f-43db-b9ff-f89e4c3f9106)

## **노드 유형**

    User: 다음 속성을 가진 사용자 표현:
        id: 사용자의 고유 식별자
        profile: 사용자 프로필 정보 (문자열로 저장)
## **관계 유형**

    FRIEND: 두 사용자 간의 양방향 관계

# **시스템 초기화 흐름**

![image](https://github.com/user-attachments/assets/959d03d3-72fe-4e81-948c-0f9bbb8dac0b)

# **기술적 구현 세부사항**

## **Neo4j 연결 및 사용**

애플리케이션은 Neo4j 그래프 데이터베이스와 상호작용하기 위해 Neo4j Go 드라이버(v5)를 사용합니다:

    연결은 database.ConnectNeo4j()에서 설정됨
    쿼리는 인젝션 공격 방지를 위해 매개변수화된 Cypher 문을 사용
    데이터베이스 작업은 ACID 준수를 위해 트랜잭션으로 래핑됨

코드베이스의 Cypher 쿼리 예시:

MERGE (u:User {id: $userID})  
ON CREATE SET u.profile = $profile  
ON MATCH SET u.profile = $profile  
RETURN u  

## **API 요청 처리**

애플리케이션은 HTTP 요청을 처리하기 위해 Gin 프레임워크를 사용합니다:

    요청은 특정 컨트롤러 메소드로 라우팅됨
    입력 유효성 검사는 BindJSON을 사용하여 컨트롤러 메소드에서 수행
    컨트롤러는 HTTP 요청을 서비스 호출로 변환
    응답은 적절한 HTTP 상태 코드와 함께 JSON으로 반환됨 userController.go:11-23

## **결론**

Friend 애플리케이션은 관심사를 분리하고 유지보수성을 촉진하는 깔끔한 계층화된 아키텍처를 따릅니다. 그래프 데이터베이스(Neo4j)의 사용은 사용자 관계를 효율적으로 표현하고 쿼리하는 소셜 네트워킹 도메인에 적합합니다. 시스템의 모듈식 설계는 전체 애플리케이션에 영향을 주지 않고 기능을 쉽게 확장하고 수정할 수 있게 합니다.

# **사용자 관리 API**

## **소개**

사용자 관리 API는 Friend 애플리케이션의 Neo4j 그래프 데이터베이스에서 사용자 노드를 생성, 수정 및 삭제하기 위한 엔드포인트를 제공합니다. 이 API는 소셜 네트워킹 서비스 내에서 사용자 신원 관리의 기초 역할을 합니다.

## **API 엔드포인트 개요**

사용자 관리 API는 세 가지 주요 엔드포인트로 구성됩니다:
|엔드포인트|	HTTP| 메소드 설명|
|-------|-------|-----|
|`/user/create`|	POST|	데이터베이스에 새 사용자 노드 생성|
|`/user/edit`|	POST|	기존 사용자의 프로필 정보 업데이트|
|`/user/delete`|	POST|	데이터베이스에서 사용자 노드 제거|

## **요청 예시**

```
POST localhost:8080/user/create  
Content-Type: application/json  
  
{  
  "userId": "test12",  
  "profile": "asdfas12"  
}
```
