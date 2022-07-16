GOPATH ?= $(HOME)/go

run:
	go run main.go

format:
	go fmt ./...

dev:
	DC_APP_ENV=dev $(GOPATH)/bin/reflex -s -r '\.go$$' make format run

test-cov:
	go test -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html