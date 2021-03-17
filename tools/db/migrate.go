package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/sudachi0114/web-db-docker/db"
	"github.com/sudachi0114/web-db-docker/models"
)

func main() {
	DBMS := "mysql"
	CONNECTION := "test:passw0rd@tcp(db:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db := db.Connect(DBMS, CONNECTION)
	defer db.Close()

	db.AutoMigrate(&models.User{})
}
