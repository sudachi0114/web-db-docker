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
