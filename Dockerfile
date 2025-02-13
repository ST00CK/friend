FROM golang:1.21-slim

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o server ./src/server.go

CMD ["/app/server"]
