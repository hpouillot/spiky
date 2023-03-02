test:
    go test -json -v spiky/pkg/models | gotestfmt

install:
    go install

run cmd:
    go run main.go {{cmd}}