package api

import (
	db "Task2DbWithGin/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowPublisher(c *gin.Context) {
	// создаем соединение с бд

	c.JSON(http.StatusOK, db.ShowPublishers())
}
