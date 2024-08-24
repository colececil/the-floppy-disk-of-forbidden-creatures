include .env

README = "BEWARE."
EULA = "THAT WHICH YOU SUMMON SHALL FOREVER BE BOUND TO YOU."

.PHONY: build build-all clean

build-all:
	@echo "Building all versions..."
	@$(MAKE) -s build OS=windows ARCH=amd64
	@$(MAKE) -s build OS=windows ARCH=arm64
	@$(MAKE) -s build OS=darwin ARCH=amd64
	@$(MAKE) -s build OS=darwin ARCH=arm64
	@$(MAKE) -s build OS=linux ARCH=amd64
	@$(MAKE) -s build OS=linux ARCH=arm64

build: clean
	@echo "Building $(OS)-$(ARCH) version..."
	@GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags '-X main.apiKey=${OPENAI_API_KEY}' -o bin/$(OS)-$(ARCH)/summon cmd/summon/summon.go
	@echo $(README) > bin/$(OS)-$(ARCH)/README.txt
	@echo $(EULA) > bin/$(OS)-$(ARCH)/EULA.txt
	@cp -r assets bin/$(OS)-$(ARCH)

clean:
	@echo "Cleaning $(OS)-$(ARCH) version..."
	@rm -rf bin/$(OS)-$(ARCH)