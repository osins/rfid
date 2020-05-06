git add .
git commit -m "build project to docker"
git push

cd docker
docker build --no-cache -t registry.haier-ioc.com/rfid-xlslr560 .
docker push registry.haier-ioc.com/rfid-xlslr560

docker tag registry.haier-ioc.com/rfid-xlslr560 wangsying/rfid-xlslr560
docker push wangsying/rfid-xlslr560

cd ..