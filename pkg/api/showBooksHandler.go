package api

import (
	"Task2DbWithGin/pkg/models"
	db "Task2DbWithGin/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowBooksHandler(c *gin.Context) {
	// Создаем структура типа Book
	var book models.Book

	// Вывод в формате json полученных данных
	c.JSON(http.StatusOK, db.GetBooks(book))

}
