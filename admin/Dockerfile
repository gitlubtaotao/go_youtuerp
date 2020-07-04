
FROM golang:1.14
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct
LABEL maintainer="Xutatao <xtt691373656@iCloud.com>"
WORKDIR /go/src/youtuerp/admin
COPY ./ ./go/src/youtuerp/admin
ENTRYPOINT go run ./main.go
