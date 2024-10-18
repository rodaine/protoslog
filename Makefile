MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory
.SUFFIXES:
SHELL := bash
.SHELLFLAGS := -euo pipefail -c
.DELETE_ON_ERROR:
.DEFAULT_GOAL := all

#: Which golangci-lint version to install for linting
GOLANGCI_LINT_VERSION = $(shell cat .versions/GOLANGCI_LINT)
#: Which golangci-lint binary to use
GOLANGCI_LINT ?= .bin/golangci-lint-$(GOLANGCI_LINT_VERSION)
#: Which buf CLI version to install for code generation
BUF_VERSION = $(shell cat .versions/BUF)
#: Which buf CLI binary to use
BUF ?= .bin/buf-$(BUF_VERSION)
#: Which protoc-gen-slog CLI to use
PGSLOG ?= .bin/protoc-gen-slog
#: Which go CLI binary to use
GO ?= go
#: Test flags for Go tests
GO_TEST_FLAGS ?= -vet=off -race -cover

codegen += internal/gen

.PHONY: all
all:
	$(MAKE) build

.PHONY: build
build: $(codegen)
	$(GO) build ./...

.PHONY: install
install: $(codegen)
	$(GO) install ./protoc-gen-slog

.PHONY: check
check: test lint

.PHONY: test
test: $(codegen)
	$(GO) test $(GO_TEST_FLAGS) ./...

.PHONY: lint
lint: $(codegen) $(GOLANGCI_LINT) $(BUF)
	$(GOLANGCI_LINT) run
	$(BUF) lint

.PHONY: generate
generate: $(BUF) $(PGSLOG)
	$(BUF) generate

.PHONY: generate-check
generate-check:
	test -z "$$(git status --porcelain | tee /dev/stderr)"

.PHONY: clean
clean:
	rm -rf $(codegen) .bin

$(codegen):
	$(MAKE) generate

$(PGSLOG): export GOBIN=$(abspath .bin)
$(PGSLOG): $(shell find ./protoc-gen-slog)
	$(GO) install ./protoc-gen-slog

$(GOLANGCI_LINT): export GOBIN=$(abspath .bin)
$(GOLANGCI_LINT):
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) \
	&& mv $(GOBIN)/golangci-lint $@

$(BUF): export GOBIN=$(abspath .bin)
$(BUF):
	$(GO) install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) \
	&& mv $(GOBIN)/buf $@