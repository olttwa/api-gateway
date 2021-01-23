SHELL := /bin/bash

.EXPORT_ALL_VARIABLES:
GO111MODULE := on

.PHONY: lint
lint:
	golint ./... | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; }

.PHONY: build
build:
	go build -o ./_output/bin/rgate ./cmd/rgate

.PHONY: install
install:
	go install ./cmd/rgate

.PHONY: test
test:
	go test -race -cover ./...