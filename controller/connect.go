package controller

import (
	"github.com/jinzhu/gorm"
)

func DBconnect() (db *gorm.DB, err error) {
	DBMS := "mysql"

	USER := "test"
	PASS := "passw0rd"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "test_db"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	// CONNECT := "test:passw0rd@tcp(localhost:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	return gorm.Open(DBMS, CONNECT)
}
