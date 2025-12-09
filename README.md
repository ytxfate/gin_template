# gin_template
gin 模板

> start
```bash
cd cmd/api 
swag init -o ../../internal/api/docs -g cmd/api/main.go -d ../../ && go run .
```

> build
```bash
bash scripts/amd_arm_build.sh -m api -v 1.0.0
```

> database container command
```bash
# GuassDB
docker run --rm --name opengauss --privileged=true -d -e GS_PASSWORD=Gaussdb@1 -e GS_USERNAME=test -e GS_USER_PASSWORD=Gaussdb@1 -p 5432:5432 opengauss:7.0.0-RC1

# MongoDB
docker run -itd --rm --name mongodb -e MONGO_INITDB_ROOT_USERNAME=test -e MONGO_INITDB_ROOT_PASSWORD=test -p 27017:27017 mongodb/mongodb-community-server:4.4-ubuntu2004 --auth

# Nacos
docker run --name nacos-standalone-derby \
    -e MODE=standalone \
    -e NACOS_AUTH_TOKEN=MTIzNDU2Nzg5MGFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6 \
    -e NACOS_AUTH_IDENTITY_KEY=MTIzNDU2Nzg5MGFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6 \
    -e NACOS_AUTH_IDENTITY_VALUE=MTIzNDU2Nzg5MGFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6 \
    -p 8080:8080 \
    -p 8848:8848 \
    -p 9848:9848 \
    -d nacos/nacos-server:v3.0.3
```
