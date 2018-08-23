FROM golang:1.10-alpine3.8
RUN apk update; apk upgrade
RUN apk add git
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/app
VOLUME ["/go/src/app"]
