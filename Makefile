.PHONY: build
build:
	go build -v ./cmd/apisecrver
.DEFAULT_GOAL := build