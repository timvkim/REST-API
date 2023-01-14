.PHONY:
.SILENT:
.DEFAULT_GOAL := all

build:
	go build -o ./.bin/app ./cmd/app/main.go

run: build
	./.bin/app