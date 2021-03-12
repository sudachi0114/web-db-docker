package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {

	// DBMS := "mysql"

	// USER := "test"
	// PASS := "passw0rd"
	// PROTOCOL := "tcp(localhost:3306)"
	// DBNAME := "test_db"
	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	// CONNECT := "test:passw0rd@tcp(localhost:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open("mysql", "test:passw0rd@tcp(localhost:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success to Connect Database.")
	defer db.Close()
}
