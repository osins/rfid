
FROM golang

COPY ./sources.aliyun.list /etc/apt/sources.list

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env

RUN go get -u -v github.com/wangsying/rfid/xlslr5603
RUN find / -name xlslr5603

RUN mkdir -p /usr/local/rfid/xlslr5603
COPY .env /usr/local/rfid/xlslr5603

RUN /go/bin/xlslr5603

WORKDIR /go/src/github.com/wangsying/rfid/xlslr5603

EXPOSE 80

CMD ["/go/src/github.com/wangsying/rfid/xlslr5603"]