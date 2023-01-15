GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell cat package.json | jq .version)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

fe-law:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/fe-law

clean:
	rm fe-law

test:
	go test -v ./...

lint:
	golangci-lint run ./...


.PHONY: \
	fe-law \
	binding \
	clean \
	test \
	lint

