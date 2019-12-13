all: test build

install:
	go install

build:
	go build -v -ldflags="-w -s" -o output/di

test:
	go test ./...
