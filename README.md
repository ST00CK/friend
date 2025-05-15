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

## **메인 애플리케이션 **

애플리케이션 진입점(main.go)은 다음과 같은 중요한 초기화 작업을 수행합니다:

    환경 설정 로드
    Neo4j 데이터베이스 연결 설정
    CORS 지원을 통한 Gin 웹 프레임워크 설정
    API 라우트 구성
    8080 포트에서 HTTP 서버 시작 

## **컨트롤러**

컨트롤러 계층은 HTTP 요청을 처리하고 이를 서비스 호출로 변환합니다. 주요 컨트롤러는 userController.go로, 다음 엔드포인트를 관리합니다:

| 엔드포인트 |	메소드 |	설명 | 코드 | 참조 |
|/user/create|	POST|	새 사용자 노드 생성|	userController.go:11-23|
|/user/delete|	POST|	사용자 노드 삭제|	userController.go:25-39|
|/user/edit|	POST|	사용자 프로필 업데이트|	userController.go:41-53|
|/friendship/create|	POST|	사용자 간 친구 관계 생성|	userController.go:55-67|
|/friendship/delete|	POST|	친구 관계 제거|	userController.go:69-81|
|/friends|	GET|	사용자의 친구 목록 조회|	userController.go:83-93|
