
FROM golang

COPY ./sources.aliyun.list /etc/apt/sources.list

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env

RUN cd /go/src/ && git clone https://github.com/wangsying/od.git
RUN ls /go/src/od
RUN cd /go/src/od && go mod vendor
RUN cd /go/src/od && go build
RUN chmod -R 777 /go/src/od/od

WORKDIR /go/src/od

EXPOSE 80

CMD ["/go/src/od/od"]