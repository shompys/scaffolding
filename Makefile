.PHONY: build run clean

build:
	go build -o bin/scaf ./cmd

run:
	go run ./cmd

clean:
	rm -rf bin/
