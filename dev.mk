.PHONY: build fmt setup

BUILDER_IMAGE := drive-build-env
BUILDER_WORK_DIR := /go/src/github.com/ty-edelweiss/crawler
BUILDER_CMD := docker run -e GOPATH=/go -v `pwd`:$(BUILDER_WORK_DIR) -w $(BUILDER_WORK_DIR) $(BUILDER_IMAGE)

build:
	$(BUILDER_CMD) make clean build

fmt:
	$(BUILDER_CMD) make fmt

setup:
	docker build -t $(BUILDER_IMAGE)