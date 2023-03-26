SHELL := /bin/bash

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell cat package.json | jq .version)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS :=-ldflags "$(LDFLAGSSTRING)"

GENPATH := "./proto/common"
COMMONPATH := "github.com/savour-labs/key-locker/proto/common"


fe-scrapy:
	env GO111MODULE=on go build $(LDFLAGS) ./cmd/fe-scrapy
.PHONY: fe-scrapy

clean:
	rm fe-scrapy

test:
	go test -v ./...

lint:
	golangci-lint run ./...


gen:
	sh ../savs/scripts/proto_gen.sh
