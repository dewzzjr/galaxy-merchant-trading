mod:
	rm -rf ./vendor
	go mod vendor -v
test:
	go test -race -cover -v ./...
build:
	go build cmd/main.go