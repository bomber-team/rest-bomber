#stage 1. Build Executable binary
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp
COPY . .

# get all dependencies
RUN go get -d -v

#build binary
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/rest_bomber

#Stage 2. Build small version of image
COPY --from=builder /go/bin/rest_bomber /go/bin/rest_bomber
ENTRYPOINT [ "/go/bin/rest_bomber" ]