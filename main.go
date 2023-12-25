package main

import (
	// "gin-tutorial/practice"
	"gin-tutorial/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	// practice.P1()

	r := gin.Default()

	r.GET("/books", handler.GetBooks)

	r.Run()

}
