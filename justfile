test:
    go test -json -v ./... | gotestfmt

install:
    go install

run cmd:
    go run main.go {{cmd}}