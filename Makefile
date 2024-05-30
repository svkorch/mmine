.PHONY: all run test build
all: test build

run:
	go run ./cmd/exchanges/main.go

test:
	go test -v -count=1 ./internal/lib/exchanger

build:
	go build -o ./bin/exchanges ./cmd/exchanges/main.go
