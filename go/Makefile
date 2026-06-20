.PHONY: build run clean

BINARY_NAME={{BINARY_NAME}}

build:
	go build -o bin/$(BINARY_NAME) ./cmd

run:
	go run ./cmd

clean:
	rm -rf bin/
