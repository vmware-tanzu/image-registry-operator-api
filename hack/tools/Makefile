# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE := on

ENVTEST_DOWNLOAD_DIR = $(shell ./$(BIN_DIR)/setup-envtest --bin-dir $(TOOLS_BIN_DIR) -p path use)

# Directories.
BIN_DIR := bin
TOOLS_BIN_DIR = $(shell pwd)/$(BIN_DIR)

HOST_OS=$(shell go env GOOS)
HOST_ARCH=$(shell go env GOARCH)

# Get the information about the platform on which the tools are built/run.
GOHOSTOS := $(shell go env GOHOSTOS)
GOHOSTARCH := $(shell go env GOHOSTARCH)

# Binaries.
CONTROLLER_GEN  := $(BIN_DIR)/controller-gen
CONVERSION_GEN  := $(BIN_DIR)/$(GOHOSTOSARCH)/conversion-gen
GOLANGCI_LINT   := $(BIN_DIR)/golangci-lint
SETUP_ENVTEST   := $(BIN_DIR)/setup-envtest

## --------------------------------------
##@ Help
## --------------------------------------

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
##@ Binaries
## --------------------------------------

.PHONY: $(CONTROLLER_GEN)
controller-gen: $(CONTROLLER_GEN) ## Install controller-gen
$(CONTROLLER_GEN): go.mod
	go build -tags=api_tools -o $@ sigs.k8s.io/controller-tools/cmd/controller-gen

.PHONY: $(GOLANGCI_LINT)
golangci-lint: $(GOLANGCI_LINT) ## Install golangci-lint
$(GOLANGCI_LINT): go.mod
	GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) \
	go build -tags=api_tools -o $(@) github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: $(SETUP_ENVTEST)
setup-envtest: $(SETUP_ENVTEST) ## Install setup-envtest
$(SETUP_ENVTEST): go.mod
	@echo "==> Installing setup-envtest"
	@go build -tags=api_tools -o $@ sigs.k8s.io/controller-runtime/tools/setup-envtest

conversion-gen: $(CONVERSION_GEN) ## Install conversion-gen
$(CONVERSION_GEN): go.mod
	GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) \
	go build -tags=api_tools -o $(@) k8s.io/code-generator/cmd/conversion-gen

## --------------------------------------
##@ Generate
## --------------------------------------

.PHONY: modules
modules: ## Runs go mod tidy to validate modules
	go mod tidy -v

.PHONY: modules-download
modules-download: ## Downloads and caches the modules
	go mod download

## --------------------------------------
##@ Cleanup
## --------------------------------------

.PHONY: clean
clean: ## Run all the clean targets
	$(MAKE) clean-bin

.PHONY: clean-bin
clean-bin: ## Remove all generated binaries
	rm -rf bin
