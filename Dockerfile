FROM golang:1.18

WORKDIR /go/src/github.com/BOOST-2021/boost-app-back
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/boost-api github.com/BOOST-2021/boost-app-back

###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/boost-api /usr/local/bin/boost-api
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["boost-api"]
