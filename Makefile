.PHONY: build dep fmt

SERVER := cmd/server

GOOS := linux
GOARCH := amd64

build: clean $(SERVER)

dep:
	dep ensure

fmt:
	go fmt ./...

% : %.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@ -ldflags "-X=main.version=$(shell git rev-parse HEAD)" $<
