package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sudachi0114/web-db-docker/db"
)

func main() {

	DBMS := "mysql"
	CONNECTION := "test:passw0rd@tcp(db:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db := db.Connect(DBMS, CONNECTION)
	defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", responseWithTemplate)
	r.GET("/ping", responseByJson)

	r.Run()
}

func responseByJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func responseWithTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
