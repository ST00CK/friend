# golang:1.23 사용
FROM golang:1.23

# 작업 디렉토리 설정
WORKDIR /app

# go.mod와 go.sum 파일 복사 후 종속성 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드 (VCS 정보 비활성화)
RUN go build -buildvcs=false -o main .

# 실행 파일 설정
CMD ["./main"]
