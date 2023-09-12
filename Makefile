build:
	go build -o bin/main ./src

start:
	./bin/main

run:
	go run ./src

compile:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 ./src
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 ./src