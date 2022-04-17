run:
	go run cmd/main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/app -ldflags="-s -w" ./cmd

compress:
	upx --best --lzma ./build/app
