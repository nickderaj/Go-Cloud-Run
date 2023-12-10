APP_NAME=goose
.DEFAULT_GOAL := build

lint:
	@go fmt ./...
	@go vet ./...
	@golangci-lint run

test:
	@go test -v

build:
	@go build -o bin/${APP_NAME}

run: build
	./bin/${APP_NAME}

clean:
	go clean
	rm ./bin/${APP_NAME}