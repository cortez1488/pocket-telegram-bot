.PHONY:
.SILENT:

build:
	go build -o .bin/bot cmd/bot/main.go

run :build
	./.bin/bot

build-container:
	docker build -t telegram-pocketbot-youtube:v0.1 .

start-container:
	docker run --name telegram-bot -p 8000:80 --env-file .env telegram-pocketbot-youtube:v0.1