#build stage
FROM golang:1.17.6-alpine AS builder
LABEL name=amaro Version=0.0.1 maintainer="Juan Farber <jfarberjob@gmail.com>"
WORKDIR /go/src/app
RUN apk add --no-cache git
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin /app

WORKDIR /app
COPY config.toml .
CMD ./cloudhumans
EXPOSE 9999