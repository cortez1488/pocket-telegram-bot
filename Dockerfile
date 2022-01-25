FROM golang:1.17.6-alpine3.15 AS builder
COPY . /github.com/cortez1488/pocket-telegram-bot/
WORKDIR /github.com/cortez1488/pocket-telegram-bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/cortez1488/pocket-telegram-bot/bin/bot .
COPY --from=0 /github.com/cortez1488/pocket-telegram-bot/configs configs/

EXPOSE 80

CMD ["./bot"]

