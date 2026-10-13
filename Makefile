
APP_NAME := anki-mcp
BUILD_DIR := ./build
GO := go

GOOS ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)

# Windows binaries require .exe
EXE_EXT := $(if $(filter windows,$(GOOS)),.exe,)

# Platform-specific output filename
BINARY := $(BUILD_DIR)/$(APP_NAME)-$(GOOS)-$(GOARCH)$(EXE_EXT)

.DEFAULT_GOAL := build

.PHONY: build build-all run test fmt vet version clean

## Build for the selected OS and architecture
build:
	@mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -trimpath -o $(BINARY) .
	@echo "Built $(BINARY)"

## Build for all supported platforms
build-all:
	$(MAKE) build GOOS=linux GOARCH=amd64
	$(MAKE) build GOOS=linux GOARCH=arm64
	$(MAKE) build GOOS=darwin GOARCH=amd64
	$(MAKE) build GOOS=darwin GOARCH=arm64
	$(MAKE) build GOOS=windows GOARCH=amd64
	$(MAKE) build GOOS=windows GOARCH=arm64

## Run locally
run:
	$(GO) run .

## Run tests
test:
	$(GO) test -v ./...

## Format code
fmt:
	$(GO) fmt ./...

## Static analysis
vet:
	$(GO) vet ./...

## Display application version
version:
	$(GO) run . --version

## Remove build artifacts
clean:
	@rm -rf $(BUILD_DIR)
