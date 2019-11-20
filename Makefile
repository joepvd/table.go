SHELL := /bin/bash

TESTS := $(shell find . -name '*_test.go')
GOFILES := $(shell find . -name '*.go')
SOURCEFILES := $(filter-out $(TESTS), $(GOFILES))
ENTRY := cmd/table.go
TARGET := table

.PHONY: build
build: $(TARGET)

$(TARGET): $(SOURCEFILES)
	go build -o $(TARGET) $(ENTRY)

.PHONY: test integration-test unit-test lint

test: integration-test unit-test lint

integration-test: build
	for f in test/*; do $$f; done

unit-test:
	go test -coverprofile=cover.out ./...

lint: .ensure-lint
	golint -set_exit_status ./...

.PHONY: .ensure-lint
.ensure-lint:
	if ! command -v go lint >/dev/null 2>&1; then go get -u golang.org/x/lint/golint; fi

.PHONY: clean
clean:
	rm -f $(TARGET) cover.out
