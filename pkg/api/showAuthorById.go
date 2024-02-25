package api

import (
	db "Task2DbWithGin/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowAuthorById(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, db.ShowAuthorById(id))

}
