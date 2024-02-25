package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	db "Task2DbWithGin/pkg/models"
)

func ShowPublisherById(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, db.ShowPublisherById(id))

}
