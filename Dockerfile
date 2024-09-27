FROM golang:1.21-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o ./cmd/bin/app ./cmd/app/main.go

CMD ["./cmd/bin/app"]