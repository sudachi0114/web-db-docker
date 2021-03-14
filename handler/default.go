package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseByJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ResponseWithTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
