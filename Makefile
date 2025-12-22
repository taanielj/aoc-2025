.PHONY: build run clean test fmt lint

BIN := bin/aoc

build:
	go build -o $(BIN) cmd/aoc/main.go

run: build
	$(BIN) $(ARGS)

run-day-%: build
	$(BIN) --day $* $(ARGS)

run-all: build
	$(BIN) --all $(ARGS)

clean:
	rm -rf bin/

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run
