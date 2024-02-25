package api

import (
	db "Task2DbWithGin/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateValueHandler(c *gin.Context) {
	//// UPDATE BY BOOKNAME
	c.JSON(http.StatusOK, db.UpdateBook(c))

}
