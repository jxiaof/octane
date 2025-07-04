FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o octane ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/octane .

CMD ["./octane"]