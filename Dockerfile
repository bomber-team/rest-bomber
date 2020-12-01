#stage 1. Build Executable binary
FROM golang:alpine as builder

WORKDIR /user/app
COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bomber.service

FROM scratch
COPY --from=builder /user/app/bomber.service /
CMD ["/bomber.service"]
