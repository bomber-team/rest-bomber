#stage 1. Build Executable binary
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp
COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/restbomber

FROM scratch

COPY --from=builder /go/bin/restbomber /go/bin/restbomber
ENTRYPOINT [ "/go/bin/restbomber" ]
