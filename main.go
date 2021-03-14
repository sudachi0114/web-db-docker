package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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
