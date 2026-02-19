.DEFAULT_GOAL := build

.PHONY:fmt vet clean mkrdir_bin build
fmt:
		go fmt ./...

vet: fmt
		go vet ./...

clean: vet
		go clean ./...

mkrdir_bin: clean
		mkdir -p bin

build: mkrdir_bin
		go build -o ./bin
