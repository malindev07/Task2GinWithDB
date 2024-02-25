package api

import (
	db "Task2DbWithGin/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddBookHadnler(c *gin.Context) {

	c.JSON(http.StatusOK, db.AddBook(c))

}
