BINARY_NAME=Vedana
MAIN_PACKAGE=./main.go

.PHONY: all build test clean

all: prepare build

prep:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest
	rm -rf ./konlpy
	git clone https://github.com/konlpy/konlpy.git
	pipx install pyinstaller
	

build:
	go mod tidy
	wails build

test:
	go test ./...
