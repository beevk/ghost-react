build:
	@go build -o bin/ghost-react # @ => run silently

run: build
	@./bin/ghost-react

test:
	@go test -v ./...

dev:
	@air