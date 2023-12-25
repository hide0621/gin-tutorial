package main

import (
	"gin-tutorial/practice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	practice.P1()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
