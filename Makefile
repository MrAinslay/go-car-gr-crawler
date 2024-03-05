run: build
	@./crawler.exe

install:
	@go mod tidy

build:
	@go build -o crawler.exe
