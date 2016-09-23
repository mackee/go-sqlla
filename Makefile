all: generate build

generate:
	go generate .

build:
	go build -v -o ./_bin/sqlla ./cmd/sqlla

.PHONY: all
