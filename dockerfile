# Bestdori-Proxy-Builder
#
# VERSION 1.0.0-rc1

FROM golang:1.20 as builder
MAINTAINER 6QHTSK <psk2019@qq.com>

ENV GO111MODULE=on

WORKDIR /go/src/Bestdori-Proxy
COPY . /go/src/Bestdori-Proxy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Bestdori-Proxy .

# Bestdori-Proxy
#
# VERSION 1.0.0-rc1
FROM alpine:latest

MAINTAINER 6QHTSK <psk2019@qq.com>

ENV GIN_MODE=release
ENV PORT=9000
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /Bestdori-Proxy

COPY --from=builder /go/src/Bestdori-Proxy/Bestdori-Proxy .

EXPOSE 9000

ENTRYPOINT ["./Bestdori-Proxy"]