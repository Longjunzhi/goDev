FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \

WORKDIR $GOPATH/src/my-api.pangxuejun.cn
COPY . $GOPATH/src/my-api.pangxuejun.cn
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./fileManage"]