.PHONY: build run clean test fmt lint

build:
	go build -o bin/aoc cmd/aoc/main.go

run: build
	./bin/aoc

clean:
	rm -rf bin/

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run
