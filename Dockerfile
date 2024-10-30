FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bot

FROM alpine:latest

RUN apk update && apk upgrade && apk add --no-cache chromium

WORKDIR /app

RUN adduser -D user
USER user

COPY --from=builder /app/bot .
COPY .env .env

ENTRYPOINT ["./bot"]

CMD ["exec", "--filter", "--hour=", "--category=", "--bash=", "--max"]
