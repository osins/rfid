### this is rfid reader event receive api server
#### start up
```
go run main.go
```

#### set databases config(root dir .env file):
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

#### docker start up
```
docker run -it wangsying/rfid-xlslr560
```

#### or exec ./build.ps1 to build your own docker container:
```
./build.ps1
```