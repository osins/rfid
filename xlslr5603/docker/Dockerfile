
FROM golang as builder

# RUN mv /etc/apt/sources.list /etc/apt/sources.list.bak
# COPY ./sources.aliyun.list /etc/apt/sources.list

# RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
# RUN go env

# RUN go get -u -v github.com/wangsying/rfid/xlslr5603
# RUN find / -name xlslr5603

RUN cd /tmp && git clone https://github.com/wangsying/rfid.git
RUN cd /tmp/rfid/xlslr5603 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o serve -ldflags="-w -s" main.go
RUN ls /tmp/rfid/xlslr5603

FROM scratch
COPY --from=builder /tmp/rfid/xlslr5603/serve /

WORKDIR /
EXPOSE 80

CMD ["/serve"]