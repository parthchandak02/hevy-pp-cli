.PHONY: build test lint install clean

build:
	go build -o bin/hevy-pp-cli ./cmd/hevy-pp-cli

test:
	go test ./...

lint:
	golangci-lint run

install:
	go install ./cmd/hevy-pp-cli

clean:
	rm -rf bin/

build-mcp:
	go build -o bin/hevy-pp-mcp ./cmd/hevy-pp-mcp

install-mcp:
	go install ./cmd/hevy-pp-mcp

build-all: build build-mcp
