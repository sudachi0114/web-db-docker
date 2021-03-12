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


### Create user

* ユーザを作成

```sql
CREATE USER 'test'@'localhost' IDENTIFIED BY 'passw0rd';
```

* 確認 (log)

```
mysql> SELECT host, user FROM mysql.user;
+-----------+---------------+
| host      | user          |
+-----------+---------------+
| localhost | mysql.session |
| localhost | mysql.sys     |
| localhost | root          |
+-----------+---------------+
3 rows in set (0.00 sec)

mysql> CREATE USER 'test'@'localhost' IDENTIFIED BY 'passw0rd';
Query OK, 0 rows affected (0.00 sec)

mysql> SELECT host, user FROM mysql.user;
+-----------+---------------+
| host      | user          |
+-----------+---------------+
| localhost | mysql.session |
| localhost | mysql.sys     |
| localhost | root          |
| localhost | test          |  # <= 作成したユーザがちゃんとできてる
+-----------+---------------+
4 rows in set (0.01 sec)
```

* 作成したユーザの削除

```sql
DROP USER test;
```


### 権限の付与

* `test` ユーザに、ローカル DB の操作に関する全ての権限を付与

```sql
GRANT all ON *.* TO 'test'@'localhost';
```

* 確認 (log)

```sql
show grants for 'test'@'localhost';
```

```
mysql> show grants for 'test'@'localhost';
+---------------------------------------------------+
| Grants for test@localhost                         |
+---------------------------------------------------+
| GRANT ALL PRIVILEGES ON *.* TO 'test'@'localhost' |
+---------------------------------------------------+
1 row in set (0.00 sec)
```


### Create database

* データベースの作成

```sql
CREATE DATABASE test_db;
```

* 確認 (log)

```
mysql -utest -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 5
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
+--------------------+
1 row in set (0.00 sec)

mysql> CREATE DATABASE test_db;
Query OK, 1 row affected (0.00 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| test_db            |
+--------------------+
5 rows in set (0.00 sec)
```


## Link
* [【第12回】Go言語（Golang）入門～MySQL接続編～](https://rightcode.co.jp/blog/information-technology/golang-introduction-mysql-connection)

* [MySQLユーザ作成について](https://qiita.com/gatapon/items/92b942fa7081cfe17482)
* [[MySQL]権限の確認と付与](https://qiita.com/shuntaro_tamura/items/2fb114b8c5d1384648aa)

* [データベースを作成する(CREATE DATABASE文)](https://www.dbonline.jp/mysql/database/index1.html)