FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/masutech16/pieceLine
COPY . .
RUN make

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/masutech16/pieceLine/app /app

ENTRYPOINT ["/app"]
