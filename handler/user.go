package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/sudachi0114/web-db-docker/models"
	"github.com/sudachi0114/web-db-docker/repo"
)

func CreateUser(c *gin.Context) {
	DBMS := "mysql"
	CONNECTION := "test:passw0rd@tcp(db:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db := repo.Connect(DBMS, CONNECTION)
	defer db.Close()

	var user models.User
	c.BindJSON(&user)

	log.Println("user", user)
	log.Println("user.Name", user.Name)
	log.Println("user.Email", user.Email)

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Accept Request",
	})
}
