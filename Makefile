.PHONY: test

all: test build

install:
	go install

build:
	go build -v -ldflags="-w -s" -o build/output/di

test:
	go test ./...
