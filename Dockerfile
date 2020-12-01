#stage 1. Build Executable binary
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o service

FROM scratch
COPY --from=builder /app/service /
EXPOSE 9999
CMD ["/service"]
