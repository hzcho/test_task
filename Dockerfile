FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/buildings/
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd/buildings/main .
COPY ./config ./config

ENTRYPOINT ["./main"]
CMD ["-c", "./config/config.yml"]

