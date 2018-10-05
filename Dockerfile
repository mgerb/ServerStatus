FROM golang:1.11.1-alpine

WORKDIR /go/src/github.com/mgerb/ServerStatus
ADD . .
RUN apk add --no-cache git alpine-sdk
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN make linux


FROM alpine:3.8

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /server-status
COPY --from=0 /go/src/github.com/mgerb/ServerStatus/dist/ServerStatus-linux .
ENTRYPOINT ./ServerStatus-linux

