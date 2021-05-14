.DEFAULT_TARGET = help

## help: Display list of commands
.PHONY: help
help: Makefile
	@sed -n 's|^##||p' $< | column -t -s ':' | sed -e 's|^| |'

## fmt: Run go fmt against code
fmt:
	go fmt ./...

## vet: Run go vet against code
vet:
	go vet ./...

## formatting: Run go vet and go fmt
formatting:
	go fmt ./...
	go vet ./...
