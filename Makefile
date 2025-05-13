# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE := on

# Directories
BIN_DIR       := bin
SAMPLES_DIR   := hack/samples
SAMPLES_BIN_DIR := $(SAMPLES_DIR)/bin
TOOLS_DIR     := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin
export PATH := $(abspath $(BIN_DIR)):$(abspath $(TOOLS_BIN_DIR)):$(abspath $SAMPLES_BIN_DIR)):$(PATH)

# Tooling binaries
CONTROLLER_GEN     := $(TOOLS_BIN_DIR)/controller-gen
CONVERSION_GEN     := $(TOOLS_BIN_DIR)/conversion-gen
GOLANGCI_LINT      := $(TOOLS_BIN_DIR)/golangci-lint
SETUP_ENVTEST      := $(TOOLS_BIN_DIR)/setup-envtest

# Allow overriding manifest generation destination directory
MANIFEST_ROOT ?= config
CRD_ROOT      ?= $(MANIFEST_ROOT)/crd/bases

# CRI_BIN is the path to the container runtime binary.
ifeq (,$(strip $(GITHUB_RUN_ID)))
# Prefer podman locally.
CRI_BIN := $(shell command -v podman 2>/dev/null || command -v docker 2>/dev/null)
else
# Prefer docker in GitHub actions.
CRI_BIN := $(shell command -v docker 2>/dev/null || command -v podman 2>/dev/null)
endif
export CRI_BIN

# The directory in which this Makefile is located. Please note this will not
# behave correctly if the path to any Makefile in the list contains any spaces.
ROOT_DIR ?= $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

# Get the GOPATH, but do not export it. This is used to determine if the project
# is in the GOPATH and if not, to use the container runtime for the
# generate-go-conversions target.
GOPATH ?= $(shell go env GOPATH)
PROJECT_SLUG := github.com/vmware-tanzu/image-registry-operator-api

# ROOT_DIR_IN_GOPATH is non-empty if ROOT_DIR is in the GOPATH.
ROOT_DIR_IN_GOPATH := $(findstring $(GOPATH)/src/$(PROJECT_SLUG),$(ROOT_DIR))

# CONVERSION_GEN_FALLBACK_MODE determines how to run the conversion-gen tool if
# this project is not in the GOPATH at the expected location. Possible values
# include "symlink" and "docker|podman".
CONVERSION_GEN_FALLBACK_MODE ?= symlink

.PHONY: all
all: lint tools generate ## Runs tests and generates all components

## --------------------------------------
##@ Help
## --------------------------------------

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
##@ Tooling
## --------------------------------------

TOOLING_BINARIES := $(CONTROLLER_GEN) $(CONVERSION_GEN) $(GOLANGCI_LINT) $(SETUP_ENVTEST)
tools: $(TOOLING_BINARIES) ## Build tooling binaries
.PHONY: $(TOOLING_BINARIES)
$(TOOLING_BINARIES):
	make -C $(TOOLS_DIR) $(@F)

## --------------------------------------
##@ Generate
## --------------------------------------

reverse = $(if $(1),$(call reverse,$(wordlist 2,$(words $(1)),$(1)))) $(firstword $(1))
GO_MOD_FILES := $(call reverse,$(shell find . -name go.mod))
GO_MOD_OP := tidy

.PHONY: $(GO_MOD_FILES)
$(GO_MOD_FILES):
	go -C $(@D) mod $(GO_MOD_OP)

.PHONY: modules
modules: $(GO_MOD_FILES)
modules: ## Validates the modules

.PHONY: modules-vendor
modules-vendor: GO_MOD_OP=vendor
modules-vendor: $(GO_MOD_FILES)
modules-vendor: ## Vendors the modules

.PHONY: modules-download
modules-download: GO_MOD_OP=download
modules-download: $(GO_MOD_FILES)
modules-download: ## Downloads and caches the modules

.PHONY: generate
generate: ## Run all code generation targets
	$(MAKE) generate-go
	$(MAKE) generate-manifests

.PHONY: generate-go
generate-go: $(CONTROLLER_GEN) ## Runs Go related generate targets
	$(CONTROLLER_GEN) \
		paths=./api/... \
		object:headerFile="$(abspath hack/boilerplate/boilerplate.go.txt)"
ifneq (0,$(GENERATE_CODE))
	go generate ./...
endif

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN) ## Generate manifests e.g. CRD, RBAC etc.
	$(CONTROLLER_GEN) \
		paths=./api/... \
		crd:crdVersions=v1 \
		output:crd:dir=$(CRD_ROOT) \
		output:none

.PHONY: generate-go-conversions
generate-go-conversions: ## Generate conversions go code

ifneq (,$(ROOT_DIR_IN_GOPATH))

# If the project is not cloned in the correct location in the GOPATH then the
# conversion-gen tool does not work. If ROOT_DIR_IN_GOPATH is non-empty, then
# the project is in the correct location for conversion-gen to work. Otherwise,
# there are two fallback modes controlled by CONVERSION_GEN_FALLBACK_MODE.

# When the CONVERSION_GEN_FALLBACK_MODE is symlink, the conversion-gen binary
# is rebuilt every time due to GNU Make, MTIME values, and symlinks. This ifeq
# statement ensures that there is not an order-only dependency on CONVERSION_GEN
# if it already exists.
ifeq (,$(strip $(wildcard $(CONVERSION_GEN))))
generate-go-conversions: $(CONVERSION_GEN)
endif

EXTRA_PEER_DIRS := 

generate-go-conversions:
	cd api && \
	$(abspath $(CONVERSION_GEN)) \
		-v 10 \
		--output-file=zz_generated.conversion.go \
		--go-header-file="$(abspath hack/boilerplate/boilerplate.go.txt)" \
		--extra-peer-dirs='$(EXTRA_PEER_DIRS)' \
		./v1alpha1 ./v1alpha2

else ifeq (symlink,$(CONVERSION_GEN_FALLBACK_MODE))

# The generate-go-conversions target uses a symlink. Step-by-step, the target:
#
# 1. Creates a temporary directory to act as a GOPATH location and stores it
#    in NEW_GOPATH.
#
# 2. Determines the path to this project under the NEW_GOPATH and stores it in
#    NEW_ROOT_DIR.
#
# 3. Creates all of the path components for NEW_ROOT_DIR.
#
# 4. Removes the last path component in NEW_ROOT_DIR so it can be recreated as
#    a symlink in the next step.
#
# 5. Creates a symlink from this project to its new location under NEW_GOPATH.
#
# 6. Changes directories into NEW_ROOT_DIR.
#
# 7. Invokes "make generate-go-conversions" from NEW_ROOT_DIR while sending in
#    the values of GOPATH and ROOT_DIR to make this Makefile think it is in the
#    NEW_GOPATH.
#
# Because make runs targets in a separate shell, it is not necessary to change
# back to the original directory.
generate-go-conversions:
	NEW_GOPATH="$$(mktemp -d)" && \
	NEW_ROOT_DIR="$${NEW_GOPATH}/src/$(PROJECT_SLUG)" && \
	mkdir -p "$${NEW_ROOT_DIR}" && \
	rm -fr "$${NEW_ROOT_DIR}" && \
	ln -s "$(ROOT_DIR)" "$${NEW_ROOT_DIR}" && \
	cd "$${NEW_ROOT_DIR}" && \
	GOPATH="$${NEW_GOPATH}" ROOT_DIR="$${NEW_ROOT_DIR}" make $@

else ifeq ($(notdir $(CRI_BIN)),$(CONVERSION_GEN_FALLBACK_MODE))

ifeq (,$(CRI_BIN))
$(error Container runtime is required for generate-go-conversions and not detected in path!)
endif

# The generate-go-conversions target will use a container runtime. Step-by-step,
# the target:
#
# 1. GOLANG_IMAGE is set to golang:YOUR_LOCAL_GO_VERSION and is the image used
#    to run make generate-go-conversions.
#
# 2. If using an arm host, the GOLANG_IMAGE is prefixed with arm64v8, which is
#    the prefix for Golang's container images for arm systems.
#
# 3. A new, temporary directory is created and its path is stored in
#    TOOLS_BIN_DIR. More on this later.
#
# 4. The flag --rm ensures that the container will be removed upon success or
#    failure, preventing orphaned containers from hanging around.
#
# 5. The first -v flag is used to bind mount the project's root directory to
#    the path /go/src/github.com/vmware-tanzu/image-registry-operator-api inside
#    of the container. This is required for the conversion-gen tool to work
#    correctly.
#
# 6. The second -v flag is used to bind mount the temporary directory stored
#    TOOLS_BIN_DIR to
#    /go/src/github.com/vmware-tanzu/image-registry-operator-api/hack/tools/bin
#    inside the container. This ensures the local host's binaries are not
#    overwritten case the local host is not Linux. Otherwise the container would
#    fail to run the binaries because they are the wrong architecture or replace
#    the binaries with Linux's elf architecture when the localhost uses
#    something else (ex. macOS is Darwin and uses mach).
#
# 7. The -w flag sets the container's working directory to where the project's
#    sources are bind mounted, /go/src/github.com/image-registry-operator-api.
#
# 8. The image calculated earlier, GOLANG_IMAGE, is specified.
#
# 9. Finally, the command "make generate-go-conversions" is specified as what
#    the container will run.
#
# Once this target completes, it will be as if the generate-go-conversions
# target was executed locally. Any necessary updates to the generated conversion
# sources will be found on the local filesystem. Use "git status" to verify the
# changes.
generate-go-conversions:
	GOLANG_IMAGE="golang:$$(go env GOVERSION | cut -c3-)"; \
	[ "$$(go env GOHOSTARCH)" = "arm64" ] && GOLANG_IMAGE="arm64v8/$${GOLANG_IMAGE}"; \
	TOOLS_BIN_DIR="$$(mktemp -d)"; \
	  $(CRI_BIN) run -it --rm \
	  -v "$(ROOT_DIR)":/go/src/$(PROJECT_SLUG) \
	  -v "$${TOOLS_BIN_DIR}":/go/src/$(PROJECT_SLUG)/hack/tools/bin \
	  -w /go/src/$(PROJECT_SLUG) \
	  "$${GOLANG_IMAGE}" \
	  make generate-go-conversions
endif

## --------------------------------------
##@ Samples
## --------------------------------------

.PHONY: list-ctrl
list-ctrl: ## Build list sample with controller client
	$(MAKE) tools
	cd $(SAMPLES_DIR); make $@

## --------------------------------------
##@ Linting
## --------------------------------------

.PHONY: lint
lint: ## Run all the lint targets
	$(MAKE) lint-go-full
	$(MAKE) lint-markdown

GOLANGCI_LINT_FLAGS ?= --fast=true
.PHONY: lint-go
lint-go: $(GOLANGCI_LINT) ## Lint codebase
	$(GOLANGCI_LINT) run -v $(GOLANGCI_LINT_FLAGS)

.PHONY: lint-go-full
lint-go-full: GOLANGCI_LINT_FLAGS = --fast=false
lint-go-full: lint-go ## Run slower linters to detect possible issues

.PHONY: lint-markdown
lint-markdown: ## Lint the project's markdown
	docker run --rm -v "$$(pwd)":/build gcr.io/cluster-api-provider-vsphere/extra/mdlint:0.17.0


## --------------------------------------
##@ Testing
## --------------------------------------

.PHONY: test
test:
test: ## Run tests
	go test -v --race ./api/...


## --------------------------------------
##@ Cleanup
## --------------------------------------

.PHONY: clean
clean: # Clean all generated or compiled files
	$(MAKE) clean-bin 
	$(MAKE) clean-crd
	$(MAKE) modules

.PHONY: clean-bin
clean-bin: ## Remove all generated binaries
	rm -rf hack/tools/bin
	rm -rf hack/samples/bin

.PHONY: clean-crd
clean-crd: ## Remove all generated crds
	rm -rf config/crd

.PHONY: verify-codegen
verify-codegen: ## Verify generated code
	hack/verify-codegen.sh
