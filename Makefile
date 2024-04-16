include .env

build:
	@echo "Building windows-amd64 version..."
	@GOOS=windows GOARCH=amd64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/windows-amd64/summon.exe cmd/summon/summon.go
	@echo "BEWARE." > bin/windows-amd64/README.txt
	@echo "THAT WHICH YOU SUMMON SHALL BE FOREVER BOUND TO YOU." > bin/windows-amd64/EULA.txt

	@echo "Building windows-arm64 version..."
	@GOOS=windows GOARCH=arm64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/windows-arm64/summon.exe cmd/summon/summon.go
	@echo "BEWARE." > bin/windows-arm64/README.txt
	@echo "THAT WHICH YOU SUMMON SHALL BE FOREVER BOUND TO YOU." > bin/windows-arm64/EULA.txt

	@echo "Building macos-amd64 version..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/macos-amd64/summon cmd/summon/summon.go
	@echo "BEWARE." > bin/macos-amd64/README.txt
	@echo "THAT WHICH YOU SUMMON SHALL BE FOREVER BOUND TO YOU." > bin/macos-amd64/EULA.txt

	@echo "Building macos-arm64 version..."
	@GOOS=darwin GOARCH=arm64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/macos-arm64/summon cmd/summon/summon.go
	@echo "BEWARE." > bin/macos-arm64/README.txt
	@echo "THAT WHICH YOU SUMMON SHALL BE FOREVER BOUND TO YOU." > bin/macos-arm64/EULA.txt

	@echo "Building linux-amd64 version..."
	@GOOS=linux GOARCH=amd64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/linux-amd64/summon cmd/summon/summon.go
	@echo "BEWARE." > bin/linux-amd64/README.txt
	@echo "THAT WHICH YOU SUMMON SHALL BE FOREVER BOUND TO YOU." > bin/linux-amd64/EULA.txt

	@echo "Building linux-arm64 version..."
	@GOOS=linux GOARCH=arm64 go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/linux-arm64/summon cmd/summon/summon.go
	@echo "BEWARE." > bin/linux-arm64/README.txt
	@echo "THAT WHICH YOU SUMMON SHALL BE FOREVER BOUND TO YOU." > bin/linux-arm64/EULA.txt
