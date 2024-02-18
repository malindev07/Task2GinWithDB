package api

import (
	"Task2DbWithGin/pkg/models"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func ShowBooksHandler(c *gin.Context) {
	// создаем соединение с бд
	conn, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(context.Background())

	var book models.Book

	// создаем запрос к бд
	rows, err := conn.Query(context.Background(), "SELECT * FROM book;")
	// выводи все строки, предварительно поместив из  в структуру
	for rows.Next() {
		rows.Scan(
			&book.BookName,
			&book.YearOfPublish,
			&book.AuthorID,
			&book.PublisherID)
		fmt.Fprintf(c.Writer,
			"Название книги: %v, ID автора: %v, год выпуска: %v, ID издателя: %v\n",
			book.BookName, book.AuthorID, book.YearOfPublish, book.AuthorID)
	}

}
