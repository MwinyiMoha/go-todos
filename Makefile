dev:
	air

run:
	go run ./cmd

build:
	CGO_ENABLED=0 GOOS=linux GOFLAGS=-mod=vendor go build -o ./build/go-todos -ldflags="-s -w" ./cmd

compress_binary:
	upx --best --lzma ./build/go-todos

test_binary:
	upx -t ./build/go-todos

test:
	go test -v -cover ./...

.PHONY: dev run build compress_binary test_binary test