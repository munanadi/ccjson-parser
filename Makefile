.PHONY: build run

build:
	@go build -o ccjsonparser

run: build
	@./ccjsonparser