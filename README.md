# go で作成した web-app と DB の連携 (with docker)

Go を用いた web アプリケーションの作成。

基本的には、バックエンド + インフラを作る練習がしたい。

docker contaienr でアプリケーションと、データベースを作成し、
それらがやりとりしてサービスが動作することを当面の目標にしておべんきょする。

<!--
## ライブラリ

* GORM (ORM); `github.com/jinzhu/gorm`
* Database driver (mysql); `github.com/go-sql-driver/mysql`
-->

## Requires

* Go
* git
* docker (docker-compose)


## アプリケーションの起動

```shell
$ docker-compose up -d

# 2回目以降?
$ docker-compose down && docker-compose up -d && docker-compose logs -f app
```

* log 

```log
$ docker-compose down && docker-compose up -d && docker-compose logs -f app
Removing web-db-docker_app_1 ... done
Removing network web-db-docker_default
Creating network "web-db-docker_default" with the default driver
Creating web-db-docker_app_1 ... done
Attaching to web-db-docker_app_1
app_1  | Hello world!  # <= とりあえずこれが出れば OK
web-db-docker_app_1 exited with code 0
```
