# WebDavD

A Go WebDav server implementation.

## Build

```
git clone https://gitee.com/autom-studio/webdavd
cd webdavd
go build -o examples/webdavd
```

## Run with Non-Daemon mode

```
cd examples
./webdavd
```

## Run with Docker & Docker Compose

```
docker build -t autom-studio/webdavd:v0.0.1 .
mkdir -p ./data/webdav
mkdir -p ./conf
cp examples/webdavd.yaml.example conf/webdavd.yaml
docker-compose up -d
```

## PS.

+ Creating directory via curl

```
# prefix path (/dir1) mush already exist
curl --user admin:123123 -XMKCOL http://192.168.15.128:8080/dir1/dir2
```

+ Uploading files via curl

```
curl --user admin:123123 -T text http://192.168.15.128:8080/dir1/
```

+ Downloading files via wget

```
wget --http-user admin --http-password 123123 http://192.168.15.128:8080/text
# quiet
wget -q --http-user admin --http-password 123123 http://192.168.15.128:8080/text -O text1
```

+ Generate bcrypted password hash string

```
./webdavd genpass -p 123123
```
