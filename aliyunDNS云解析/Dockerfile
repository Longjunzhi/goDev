FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \
    ENV=prod

WORKDIR $GOPATH/src/aliyundns
COPY . $GOPATH/src/aliyundns
RUN go build .
ENTRYPOINT ["./dns"]