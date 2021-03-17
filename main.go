package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/sudachi0114/web-db-docker/handler"
	"github.com/sudachi0114/web-db-docker/repo"
)

func main() {

	DBMS := "mysql"
	CONNECTION := "test:passw0rd@tcp(db:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db := repo.Connect(DBMS, CONNECTION)
	defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", handler.ResponseWithTemplate)
	r.GET("/ping", handler.ResponseByJson)

	r.GET("/new_user", tmpResponce)

	r.Run()
}

func tmpResponce(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not Implemented",
	})
}
