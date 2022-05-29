FROM golang:1.18.2-buster@sha256:7b7918c3622495e7da7f070dfaf332b19de4d6394503763e3a9272eb39e7a08a as builder
ENV GO111MODULE=on
WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o micro-api-gateway

FROM alpine:latest
ENV TZ=Asia/Ho_Chi_Minh
RUN apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/app/micro-api-gateway /app/micro-api-gateway
ENTRYPOINT ./micro-api-gateway