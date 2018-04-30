SOURCES ?= $(shell find . -path "./vendor" -prune -o -type f -name "*.go" -print)

all: init build

build: $(SOURCES)
		go build -o app

.PHONY: init
init:
		go get -u github.com/golang/dep/cmd/dep
		dep ensure