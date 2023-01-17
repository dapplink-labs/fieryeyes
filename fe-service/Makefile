GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell cat package.json | jq .version)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

fe-service:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/fe-service

clean:
	rm fe-service

test:
	go test -v ./...

lint:
	golangci-lint run ./...


.PHONY: \
	fe-service \
	binding \
	clean \
	test \
	lint

