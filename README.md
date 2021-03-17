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
$ make compose/up
```

* access check

```shell
$ curl http://localhost:8080/ping
{"message":"pong"}
```

## データベース

環境変数に入れ込めば、勝手に mysql_user を作ってくれたり、database を作ってくれたりするらしい。

```yml
version: '3'

services: 
  app:
    # (略)
  db:
    image: mysql:5.7
    environment:                # ここでいろいろ設定しておけば、いい感じに作ってくれる
      MYSQL_ROOT_PASSWORD: root
      MYSQL_USER: test
      MYSQL_PASSWORD: passw0rd
      MYSQL_DATABASE: test_db 
      TZ: 'Asia/Tokyo'
    volumes: 
      - data:/var/lib/mysql
      - ./db/my.cnf:/etc/mysql/conf.d/my.cnf
    ports: 
      - 3306:3306

volumes: 
  data:
    driver: local
```

* 確認 log

```
$ docker-compose exec db bash

root@d10a809d00ec:/# 
root@d10a809d00ec:/# mysql -utest -p 
Enter password:    # <= passw0rd
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 3
Server version: 5.7.32 MySQL Community Server (GPL)

Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| test_db            |
+--------------------+
2 rows in set (0.00 sec)

mysql> 
mysql> use test_db;
Database changed
mysql> show tables;
Empty set (0.00 sec)

mysql> 
mysql> 
mysql> exit
Bye
root@d10a809d00ec:/# 
root@d10a809d00ec:/# exit
exit
```

* app と db の (簡易的な) 接続チェック

```
$ make compose/up

$ docker-compose logs -f app
Attaching to web-db-docker_app_1
app_1  | # ... (途中 略)
app_1  | Successed to connect MySQL Database!  # <= これが出ていれば OK!
app_1  | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
app_1  | 
app_1  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
app_1  |  - using env:  export GIN_MODE=release
app_1  |  - using code: gin.SetMode(gin.ReleaseMode)
app_1  | 
app_1  | [GIN-debug] Loaded HTML Templates (2): 
app_1  |        - 
app_1  |        - index.html
app_1  | 
app_1  | [GIN-debug] GET    /                         --> main.responseWithTemplate (3 handlers)
app_1  | [GIN-debug] GET    /ping                     --> main.responseByJson (3 handlers)
app_1  | [GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
app_1  | [GIN-debug] Listening and serving HTTP on :8080
```

## エンドポイント

### テスト用

* `/` : `method:GET`

```shell
$ curl localhost:8080/
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>web-db-docker</title>
</head>
<body>
    <h1>Hello world!</h1>
</body>
</html>
```

* `/ping` : `method:GET`

```shell
$ curl localhost:8080/ping
{"message":"pong"}
```

### アプリケーションサービスとしてのエンドポイント

* `/new_user` : `method:GET`

```shell
$ curl -X POST -H "Content-Type: application/json" -d '{"name":"sudachi", "email":"sudachi@example.com"}' localhost:8080/new_user

```


## Links
* [Go+MySQL+Dockerで簡単なCRUDを実装する](https://qiita.com/daiki-murakami/items/c8f9df8defc937e185ee)
* [【第12回】Go言語（Golang）入門～MySQL接続編～](https://rightcode.co.jp/blog/information-technology/golang-introduction-mysql-connection)

* [docker go image](https://hub.docker.com/_/golang)

* [GOのORMを分かりやすくまとめてみた【GORM公式ドキュメントの焼き回し】](https://qiita.com/gold-kou/items/45a95d61d253184b0f33#select)

* [Go 構造体を JSON に変換する方法](https://www.delftstack.com/ja/howto/go/how-to-convert-go-struct-to-json/)
* [Package json](https://golang.org/pkg/encoding/json/#Marshal)

* [go-ginでサクッとRESTAPIを構築する](https://qiita.com/shiei_kawa/items/eddf48287455380f618f)