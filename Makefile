tidy:
	go mod tidy
mod:
	rm -rf ./vendor
	go mod vendor -v
test:
	go vet ./...
	go test -race -cover -v ./...
build:
	go build cmd/main.go

all: tidy mod test build