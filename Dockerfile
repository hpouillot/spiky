FROM golang:1.20-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/spiky

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/spiky .

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

WORKDIR /app

COPY ./mnist /app/mnist

COPY --from=build_base /tmp/spiky/out/spiky /app/spiky

# Run the binary program produced by `go install`
CMD ["/app/spiky", "train"]