cd ..

git add .
git commit -m "build project to docker"
git push

docker run -it -v $PWD/:/go/src/github.com/wangsying/rfid/xlslr5603 -it golang go build -o /go/src/github.com/wangsying/rfid/xlslr5603/server /go/src/github.com/wangsying/rfid/xlslr5603/main.go
Copy-Item server docker/

cd xlslr5603/docker
docker build --no-cache -t registry.haier-ioc.com/rfid-xlslr560 .
# docker push registry.haier-ioc.com/rfid-xlslr560
# docker tag registry.haier-ioc.com/rfid-xlslr560 wangsying/rfid-xlslr560
# docker push wangsying/rfid-xlslr560

cd ..