FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \
    ENV=prod

WORKDIR $GOPATH/src/aliyunDns
COPY . $GOPATH/src/aliyunDns
RUN go build .
ENTRYPOINT ["./aliyunDns"]