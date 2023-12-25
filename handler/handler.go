package handler

import (
	"gin-tutorial/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Books)
}
