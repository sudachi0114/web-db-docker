package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success to Connect Database.")
	defer db.Close()
}
