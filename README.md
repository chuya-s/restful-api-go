# vscode

## Language Server
Code > Preferences > Settings
check `go.useLanguageServer`

## GO111MODULE
setting.json in vscode
(cmd+shift+P with command pallet)
```
 "go.toolsEnvVars": {"GO111MODULE": "on"},
```

## gopls
setting.json in vscode
(cmd+shift+P with command pallet)
check `gopls: Language Server from Google`

## Reference
https://qiita.com/chanhama/items/a21ca7d5cd43d6f3f90d

# Rest Architecture
## JSON
encodeing/json
## Marshalling
data structures in GO into JSON
## Getting Started With A Basic API

# go mod
## go mod init
`$ go mod init example.com/XXX`
## go get
`$ go get github.com/gorilla/mux`


# Docker (Local)
## Dockerfile
edit Dockerfile
## docker build
`$ docker build -t backend .`
## docker run
`$ docker run -d -it -p 8080:8080 backend`

## DB
### mysql
`$ cd mysql`

`$ docker compose up -d`
```
...
[+] Running 3/3
 ⠿ Network mysql_default Created 0.0s
 ⠿ Volume "mysql-volume"  Created 0.0s
 ⠿ Container db           Started
```
`$ docker compose ps`
```
NAME                COMMAND                  SERVICE             STATUS              PORTS
db                  "docker-entrypoint.s…"   db                  running             0.0.0.0:3306->3306/tcp
```
`$ docker exec -it db bash`

`bash-4.4# mysql -utest_user -ppassword test_database`

`mysql>show tables;`
```
+-------------------------+
| Tables_in_test_database |
+-------------------------+
| article                 |
+-------------------------+
1 row in set (0.01 sec)
```
`mysql> select * from article;`
```
+----+---------+------------------+
| id | title   | content          |
+----+---------+------------------+
|  1 | 記事1   | 記事1です。      |
|  2 | 記事2   | 記事2です。      |
+----+---------+------------------+
2 rows in set (0.00 sec)
```

if remove volume of docker

`$ docker volume rm mysql-volume`

confirm volume

`$ docker volume ls`
# makefile
## make run
rest api with go by docker run 

## make ps
docker ps
