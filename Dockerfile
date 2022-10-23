FROM golang:1.18 AS Build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/minidelivery cmd/server/main.go
