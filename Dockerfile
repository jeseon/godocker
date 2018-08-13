FROM golang:alpine AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
ADD . /go/src/godocker
WORKDIR /go/src/godocker

# Install tools required for project
RUN apk add --no-cache git dep && rm -rf /var/cache/apk/*

# Install library dependencies
RUN dep ensure -vendor-only

# Go static analyzer
RUN go vet

# Run all the tests with the race detector enabled
RUN go test -v -cover

# Rebuild when a file changes in the project directory
RUN go build -a -o /go/bin/godocker

## This results in a single layer image
FROM scratch
COPY --from=builder /go/bin/godocker /godocker
EXPOSE 5000
CMD ["/godocker"]