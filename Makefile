SHELL := /bin/bash

TESTS := $(shell find . -name '*_test.go')
GOFILES := $(shell find . -name '*.go')
SOURCEFILES := $(filter-out $(TESTS), $(GOFILES))
ENTRY := cmd/table.go
TARGET := table

.PHONY: all
all: clean build test

.PHONY: build
build: $(TARGET)

$(TARGET): $(SOURCEFILES)
	go build -o $(TARGET) $(ENTRY)

.PHONY: test integration-test unit-test lint

test: integration-test unit-test lint

integration-test: build .ensure-bats .ensure-git
	test/end-to-end

unit-test: .ensure-go $(GOFILES)
	go test -v -coverprofile=cover.out ./...

.PHONY: unit
unit: unit-test

lint: .ensure-golint .ensure-golangci-lint
	golint -set_exit_status ./...
	golangci-lint run

cover.out: unit-test

.PHONY: cover-report
cover-report: cover.out
	go tool cover -html=cover.out

.PHONY: .ensure-bats .ensure-golint .ensure-go .ensure-git .ensure-golangci-lint
.ensure-%:
	@command -v $* >/dev/null 2>&1 || echo Missing $*

.PHONY: clean
clean:
	rm -f $(TARGET) cover.out
