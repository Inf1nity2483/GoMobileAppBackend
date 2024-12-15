FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod  ./
COPY go.sum ./
COPY migrations ./migrations

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/migrations ./migrations

EXPOSE 8000

CMD ["./main"]