# ローカルの MySQL とやりとり (with GORM)

Go を用いた DB 連携の練習。

ローカルで動いている MySQL に接続してやりとりしてみる。

## ライブラリ

* GORM (ORM); `github.com/jinzhu/gorm`
* Database driver (sqlite); `github.com/mattn/go-sqlite3`

## Requires

* Go
* MySQL


### Install MySQL@5,7

* Install in Mac with brew.

```shell
$ brew install mysql@5.7
$ echo 'export PATH="/usr/local/opt/mysql@5.7/bin:$PATH"' >> /Users/sudachi/.bashrc

# Install CLI client
$ brew install mysql-client@5.7
$ echo 'export PATH="/usr/local/opt/mysql-client@5.7/bin:$PATH"' >> /Users/sudachi/.bashrc
```

個人的な事情ですが、PATH の設定などは全て `.bashrc` で管理しているので、ここに登録しました。
デフォルトでは `.barh_profile` となっています。

```shell
# 参考
$ echo 'export PATH="/usr/local/opt/mysql@5.7/bin:$PATH"' >> /Users/sudachi/.bash_profile
$ echo 'export PATH="/usr/local/opt/mysql-client@5.7/bin:$PATH"' >> /Users/sudachi/.bash_profile
```


### Connect to MySQL with CLI client

```shell
# まずはサービスを起動
$ mysql.server start
Starting MySQL
. SUCCESS! 

# CLI client を用いて接続
$ mysql -uroot -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 2
Server version: 5.7.32 Homebrew

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
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.01 sec)

mysql> exit
Bye

# サービスを停止するには
$ mysql.server stop
Shutting down MySQL
. SUCCESS! 
```


## Link
* [【第12回】Go言語（Golang）入門～MySQL接続編～](https://rightcode.co.jp/blog/information-technology/golang-introduction-mysql-connection)