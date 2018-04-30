SOURCES ?= $(shell find . -path "./vendor" -prune -o -type f -name "*.go" -print)

pieceline: $(SOURCES)
		go build

.PHONY: init
init:
		go get -u github.com/golang/dep/cmd/dep
		dep ensure