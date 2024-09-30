FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/buildings/

RUN go build -o main .

FROM golang:latest

WORKDIR /app

COPY --from=builder /app/cmd/buildings/main .

COPY ./config ./config

ENTRYPOINT ["./main"]

CMD ["-c", "./config/config.yml"]
