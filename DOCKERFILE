FROM golang:1.18 AS Build

WORKDIR /app

COPY . .

RUN go mody tidy --compact= 1.18

RUN build -o /app/minidelivery cmd/server/main.go

CMD ["/app/minidelivery"]