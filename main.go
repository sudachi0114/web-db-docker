package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sudachi0114/web-db-docker/controller"
)

func main() {
	db, err := controller.DBconnect()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success to Connect Database.")
	defer db.Close()
}
