package api

import (
	"Task2DbWithGin/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"context"
	"log"
)

func AddBookHadnler(c *gin.Context) {

	conn, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(context.Background())

	var book models.Book

	// биндим занчения из post  в json формат и передаем в структуру book
	if c.BindJSON(&book) != nil {
		fmt.Fprint(c.Writer, "Возникла ошибка")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	// тело запроса к бд
	sqlReq := "INSERT INTO book (bookname, yearofpublish, author_id, publisher_id) VALUES ($1,$2,$3,$4)"

	// отправка данных в бд
	if _, err := conn.Exec(context.Background(), sqlReq,
		book.BookName,
		book.YearOfPublish,
		book.AuthorID,
		book.PublisherID); err != nil {
		fmt.Fprint(c.Writer, "Error\n")
	} else {
		fmt.Fprint(c.Writer, "Книга добавлена!\n")
	}
	_ = "\n"
	c.JSON(http.StatusCreated, book)

}
