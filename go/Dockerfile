FROM golang:1.12.6-alpine3.10

WORKDIR /go

COPY src/wserver /go/src/wserver

RUN go build -ldflags "-w" -o /go/bin/wserver /go/src/wserver/wserver.go

EXPOSE 8000

ENTRYPOINT ["/go/bin/wserver"]
