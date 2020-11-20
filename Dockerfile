FROM golang:1.15-alpine

WORKDIR /go/src/go-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get -u golang.org/x/lint/golint
