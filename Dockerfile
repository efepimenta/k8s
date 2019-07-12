FROM golang:1.12.6-alpine3.10

RUN apk update \
    && apk add --virtual build-dependencies build-base gcc

WORKDIR /go

COPY go/src/wserver /go/src/wserver

RUN go build -ldflags "-w" -o /go/bin/wserver /go/src/wserver/wserver.go

RUN apk del build-dependencies build-base gcc

#EXPOSE 8000

ENTRYPOINT ["go"]
