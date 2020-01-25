VERSION := $(shell git describe --tags)

_bin/sqlla: *.go
	go generate
	go build -o _bin/sqlla -ldflags="-X main.Version=$(VERSION)" cmd/sqlla/main.go

.PHONY: clean install get-deps test build

test: generate
	go test -v -race ./...
	cd _example && go test -v -race ./...
	go vet ./...
	cd _example && go vet ./...

clean:
	rm -Rf _bin/* _artifacts/*

install: _bin/sqlla
	install _bin/sqlla $(GOPATH)/bin

generate:
	go generate

build: clean generate test
	go generate
	mkdir -p _artifacts
	goxz -pv=${VERSION} -d=_artifacts -build-ldflags="-w -s -X main.Version=$(VERSION)" ./cmd/sqlla

release:
	ghr ${VERSION} _artifacts
