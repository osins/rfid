
FROM golang as builder

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env -w GOOS="linux"
RUN go env -w CGO_LDFLAGS="-w -s"
RUN go env

RUN cd /tmp && git clone https://github.com/wangsying/rfid.git
RUN cd /tmp/rfid/xlslr5603 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o serve main.go
RUN ls /tmp/rfid/xlslr5603

FROM scratch
COPY --from=builder /tmp/rfid/xlslr5603/serve /

WORKDIR /
EXPOSE 80

CMD ["/serve"]