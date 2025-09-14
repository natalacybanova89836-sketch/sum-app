build:
	go build -o app .

run:
	go run main.go

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

test:
	go test -v ./...

vet:
	go vet ./...
