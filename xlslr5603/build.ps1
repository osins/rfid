cd ..

git add .
git commit -m "build project to docker"
git push

cd xlslr5603/docker
docker build --no-cache -t registry.haier-ioc.com/rfid-xlslr560 .
docker push registry.haier-ioc.com/rfid-xlslr560

ssh -p 2022 root@huawei

cd ..