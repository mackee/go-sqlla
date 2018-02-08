VERSION := $(shell git describe --tags)

_bin/sqlla: *.go
	go generate
	go build -o _bin/sqlla -ldflags="-X main.Version=$(VERSION)" cmd/sqlla/main.go

.PHONY: clean install get-deps test build

test:
	go test -v -race
	go vet

get-deps:
	go get github.com/golang/dep/cmd/dep
	dep ensure

clean:
	rm -Rf _bin/* _artifacts/*

install: _bin/sqlla
	install _bin/sqlla $(GOPATH)/bin

build: clean get-deps test
	go generate
	mkdir -p _artifacts
	goxz -pv=${VERSION} -d=_artifacts -build-ldflags="-w -s -X main.Version=$(VERSION)" ./cmd/sqlla

release:
	ghr ${VERSION} _artifacts
