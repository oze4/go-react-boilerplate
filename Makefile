build:
	go build -o bin/main ./cmd/main.go

run-build:
	./bin/main

start:
	make build && make run-build

run:
	go run ./cmd/main.go