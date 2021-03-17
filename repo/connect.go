package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Connect(DBMS, CONNECTION string) *gorm.DB {
	db, err := gorm.Open(DBMS, CONNECTION)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successed to connect MySQL Database!")
	return db
}
