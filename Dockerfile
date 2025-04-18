FROM golang:1.24

WORKDIR /app

COPY go_server/ ./

RUN go mod download 

CMD ["go", "run", "/app/main.go"]