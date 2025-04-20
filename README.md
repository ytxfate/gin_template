# gin_template
gin 模板


> database container command
```bash
# GuassDB
docker run --rm --name opengauss --privileged=true -d -e GS_PASSWORD=Gaussdb@1 -e GS_USERNAME=ytx -e GS_USER_PASSWORD=Gaussdb@1 -p 5432:5432 opengauss:7.0.0-RC1

# MongoDB
docker run -itd --name mongodb -e MONGO_INITDB_ROOT_USERNAME=test -e MONGO_INITDB_ROOT_PASSWORD=test -p 27017:27017 mongodb/mongodb-community-server:4.4-ubuntu2004 --auth
```
