# Step 1: 빌드 단계
FROM golang:1.23 AS builder

# 컨테이너 내 작업 디렉토리 설정
WORKDIR /app

# 의존성 파일 복사 및 설치
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 실행 파일 생성
RUN go build -o main .

# Step 2: 실행 단계
FROM debian:bullseye-slim

# 컨테이너 내 작업 디렉토리 설정
WORKDIR /app

# 빌드 단계에서 생성한 실행 파일 복사
COPY --from=builder /app/main .

# 앱이 사용할 포트 설정 (예: 8080)
EXPOSE 8080

# 컨테이너 시작 시 실행될 명령어 설정
CMD ["./main"]
