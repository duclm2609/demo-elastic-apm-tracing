FROM golang as builder
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o micro-api-gateway

FROM alpine:latest
ENV TZ=Asia/Ho_Chi_Minh
RUN apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/app/micro-review /app/
ENTRYPOINT ./micro-review