build:
	go build -o bin/ ./cmd/...

test:
	go test -v ./...

run:
	go run ./cmd/...

start:
	./bin/cmd
.PHONY: build test
