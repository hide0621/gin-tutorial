package handler

import (
	"errors"
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

	models.Books = append(models.Books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)

}

func BookById(c *gin.Context) {

	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, book)

}

func getBookById(id string) (*models.Book, error) {

	for i, b := range models.Books {
		if b.ID == id {
			return &models.Books[i], nil
		}
	}

	return nil, errors.New("book not found")

}
