test:
    go test -json ./pkg/... | gotestfmt

install:
    go install

run cmd:
    go run main.go {{cmd}}