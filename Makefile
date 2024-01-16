build:
	@go build -o bin/tried-passwords

run: build
	./bin/tried-passwords