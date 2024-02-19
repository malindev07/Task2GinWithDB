package main

import (
	"Task2DbWithGin/pkg/api"
	"github.com/gin-gonic/gin"
)

var DbUrl = "postgres://postgres:Respekt42@localhost:5432/BookShop"

func main() {
	//соединение с базой

	r := gin.Default()
	// Вся таблица
	r.GET("/api/book", api.ShowBooksHandler)
	// ID и имена авторов
	r.GET("/api/name", api.ShowNameAndIdHandler)
	// ID и издатель
	r.GET("/api/publisher", api.ShowPublisher)
	// Имя автора по id
	r.GET("/api/author/:id", api.ShowAuthorById)
	// Имя издателя по id
	r.GET("/api/publisher/:id", api.ShowPublisherById)
	// Добавить книгу в таблицу
	r.POST("/api/addBook", api.AddBookHadnler)
	// Обновление данных
	r.PUT("/api/updateValue", api.UpdateValueHandler)
	// Удаление записи
	r.DELETE("/api/deleteValue/:bookName", api.DeleteValueHandler)
	//Слушаем сервер
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
