FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bot

FROM alpine:latest

RUN apk add chromium 

RUN adduser -D user
USER user

WORKDIR /app

COPY --from=builder /app/bot .

CMD ["./bot", "exec"]
