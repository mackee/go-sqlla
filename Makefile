all: generate build

generate:
	go get github.com/jessevdk/go-assets-builder
	go generate .

build:
	go build -v -o ./_bin/sqlla ./cmd/sqlla

.PHONY: all
