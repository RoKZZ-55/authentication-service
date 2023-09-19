# syntax=docker/dockerfile:1
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/authentication-service/
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/authentication-service ./cmd/app


FROM alpine
COPY --from=0 /go/bin/authentication-service /authentication-service/authentication-service
COPY /config /authentication-service/config
EXPOSE 8080
ENTRYPOINT ["/authentication-service/authentication-service"]