VERSION := $(shell git describe --tags)

_bin/sqlla: *.go
	go generate
	go build -o _bin/sqlla cmd/sqlla/main.go -ldflags="-X main.Version=$(VERSION)"

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
	gox -output "_artifacts/{{.Dir}}-{{.OS}}-{{.Arch}}-${VERSION}/sqlla" -ldflags "-w -s -X main.Version=$(VERSION)"
	cd _artifacts/ && find . -name 'sqlla*' -type d | sed 's/\.\///' | xargs -I{} zip -m -q -r {}.zip {}

release:
	ghr ${VERSION} _artifacts
