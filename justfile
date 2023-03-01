test:
    go test -json -v spiky/pkg/core spiky/pkg/codec | gotestfmt

install:
    go install

run cmd:
    go run main.go {{cmd}}