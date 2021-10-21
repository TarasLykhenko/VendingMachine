GOOS=linux
BUILD_OUTPUT_BINARY=vendingmachine

ifndef BUILD_OUTPUT
    BUILD_OUTPUT=./bin
endif


# Manual build
LDFLAGS=-s -w


ifeq ($(OS),Windows_NT)
	BUILD_OUTPUT_BINARY=manager.exe
	GOOS=windows
endif
ifeq ($(OS),OSX)
	GOOS=darwin
endif

.DEFAULT_GOAL := default

build:
	env GOOS=$(GOOS) go build -ldflags="$(LDFLAGS)" -o $(BUILD_OUTPUT)/$(BUILD_OUTPUT_BINARY) main.go


default: build

.PHONY: all