FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/masutech16/pieceLine
COPY . .
RUN make

# runtime image
FROM alpine
ENV TWITTER_CONSUMER_KEY=GGYCSYsMKfb606qwBKYdQjt0N
ENV TWITTER_CONSUMER_KEY_SECRET=qBJza0CPGgipNLwL58f6nY9QEDni9dAuXzK4kMm3skPq9OyxiJ
ENV TWITTER_ACCESS_TOKEN=707251965613322240-8OuaLN9v48PqrWSrQetMPDNYP48cb2g
ENV TWITTER_ACCESS_TOKEN_SECRET=irIq33ITD6IsZDrb5TvqBZqxnquCBgstcaBaUviR0jRQm
ENV NAME=masutech
ENV PASS=12qwaszx
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/masutech16/pieceLine/app /app

ENTRYPOINT ["/app"]
