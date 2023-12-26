package handler

import (
	"gin-tutorial/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Books)
}

func CreateBooks(c *gin.Context) {

	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books := append(models.Books, newBook)

	c.IndentedJSON(http.StatusCreated, books)

}
