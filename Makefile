include .env

build:
	@echo "Building windows-amd64 version..."
	@GOOS=windows GOARCH=amd64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/windows-amd64/summon.exe cmd/summon/summon.go