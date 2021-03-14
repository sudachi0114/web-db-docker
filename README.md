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

$ curl http://localhost:8080/
<!DOCTYPE html>
<html lang="ja">
<head>
    <!-- (略) -->
</head>
<body>
    <h1>Hello world!</h1>
</body>
</html>
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