VERSION := $(shell git describe --tags)

export PATH=$(shell echo $$PWD/_gobin):$(shell echo $$PATH)
export GOBIN=${PWD}/_gobin

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

get-deps:
	go mod download
	cd _example && go mod download
	mkdir -p _gobin
	go install github.com/Songmu/goxz/cmd/goxz@latest
	go install github.com/tcnksm/ghr@latest
	go install github.com/mackee/go-genddl/cmd/genddl@latest

generate: get-deps
	go generate ./...

build: clean test
	mkdir -p _artifacts
	goxz -pv=${VERSION} -d=_artifacts -build-ldflags="-w -s -X main.Version=$(VERSION)" ./cmd/sqlla

release: get-deps
	ghr ${VERSION} _artifacts
