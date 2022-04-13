.PHONY:
.SILENT:


build:
	go build -o ./.bin/wsChat cmd/main.go

run: build
	./.bin/wsChat