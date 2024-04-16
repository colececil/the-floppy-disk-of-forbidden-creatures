include .env

build:
	@echo "Building windows-amd64 version..."
	@GOOS=windows GOARCH=amd64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/windows-amd64/summon.exe cmd/summon/summon.go

	@echo "Building windows-arm64 version..."
	@GOOS=windows GOARCH=arm64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/windows-arm64/summon.exe cmd/summon/summon.go

	@echo "Building macos-amd64 version..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/macos-amd64/summon cmd/summon/summon.go

	@echo "Building macos-arm64 version..."
	@GOOS=darwin GOARCH=arm64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/macos-arm64/summon cmd/summon/summon.go

	@echo "Building linux-amd64 version..."
	@GOOS=linux GOARCH=amd64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/linux-amd64/summon cmd/summon/summon.go

	@echo "Building linux-arm64 version..."
	@GOOS=linux GOARCH=arm64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/linux-arm64/summon cmd/summon/summon.go
