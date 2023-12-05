VERSION 0.7
FROM golang:1.20
WORKDIR /app

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    RUN go install github.com/jstemmer/go-junit-report/v2@latest
    RUN go install github.com/jstemmer/go-junit-report/v2@latest
    # Output these back in case go mod download changes them.
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

test:
    FROM +deps
    COPY . .
    RUN go test -v 2>&1 ./... | go-junit-report -set-exit-code > ./junit-results.xml
    SAVE ARTIFACT junit-results.xml AS LOCAL junit-results.xml

test:
    FROM +deps
    COPY . .
    RUN go test -race -coverprofile=coverage.out -covermode=atomic
    SAVE ARTIFACT coverage.out AS LOCAL coverage.out

build:
    FROM +deps
    COPY . .
    RUN goreleaser release --snapshot --clean