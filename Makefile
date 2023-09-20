include .env

APP_NAME := go-boilerplate
SOURCE_PATH := ./src/

DBMATE_URL := postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

.PHONY: build
build:
	go build -v -o bin/${APP_NAME} ./src

.PHONY: start
start:
	./bin/${APP_NAME}

.PHONY: run
run:
	go run ./src

.PHONY: compile
compile:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 ./src
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 ./src

.PHONY: migration-up
migration-up:
	migrate -database ${DBMATE_URL} -path migrations up

.PHONY: migration-down
migration-down:
	migrate -database ${DBMATE_URL} -path migrations down