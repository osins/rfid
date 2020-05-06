### This is a receive api server by rfid reader event
#### Start up
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

#### Docker start up
```
docker run -it wangsying/rfid-xlslr560
```

#### Or exec ./build.ps1 to build your own docker container:
```
./build.ps1
```