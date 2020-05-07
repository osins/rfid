### This is a receive api server by xlslr5603 rfid reader event
#### Start up:
```
go run main.go
```

#### Set databases config(root dir .env file):
```
.env

LISTEN_HOST=
LISTEN_PORT=80

DB_HOST=192.168.5.35
DB_PORT=3306
DB_USERNAME=xlslr5603
DB_PASSWORD=xlslr5603
DB_NAME=xlslr5603
DB_TABLE_PREFIX=slr5603_
```

#### Docker start up:
```
docker run -it wangsying/rfid-xlslr560
```

#### Or exec ./build.ps1 to build your own docker container:
```
./build.ps1
```

#### About build smallest docker container:
```

FROM golang as builder

RUN cd /tmp && git clone https://github.com/wangsying/rfid.git
RUN cd /tmp/rfid/xlslr5603 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o serve -ldflags="-w -s" main.go

FROM scratch
COPY --from=builder /tmp/rfid/xlslr5603/serve /

WORKDIR /
EXPOSE 80

CMD ["/serve"]
```