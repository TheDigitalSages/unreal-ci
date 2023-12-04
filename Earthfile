VERSION 0.7
FROM golang:1.20
WORKDIR /app

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    # Output these back in case go mod download changes them.
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

test:
    FROM +deps
    COPY . .
    RUN go test -v 2>&1 ./... | go-junit-report -set-exit-code > ./junit-results.xml
    SAVE ARTIFACT junit-results.xml AS LOCAL junit-results.xml

build:
    FROM +deps
    COPY . .
    RUN go install github.com/jstemmer/go-junit-report/v2@latest
    RUN goreleaser release --snapshot --clean