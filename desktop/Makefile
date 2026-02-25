BINARY_NAME=Vedana
MAIN_PACKAGE=./main.go

.PHONY: all build test clean

all: build

build:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest
	go mod tidy
	wails build

test:
	go test ./...
