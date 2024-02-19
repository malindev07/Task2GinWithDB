package api

import (
	"Task2DbWithGin/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"context"
	"log"
)

type OldNewValue struct {
	OldValue string `json:"oldValue"`
	NewValue string `json:"newValue"`
}

func UpdateValueHandler(c *gin.Context) {
	// UPDATE BY BOOKNAME

	conn, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(context.Background())

	var book models.Book
	var update OldNewValue

	if c.BindJSON(&update) != nil {
		fmt.Fprint(c.Writer, "Возникла ошибка")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	sqlData := fmt.Sprintf("'%v'", update.OldValue)

	sqlReq := "SELECT * FROM book WHERE bookname =" + sqlData

	err = conn.QueryRow(context.Background(), sqlReq).Scan(
		&book.BookName,
		&book.YearOfPublish,
		&book.AuthorID,
		&book.PublisherID)

	fmt.Fprint(c.Writer, book, "\n")

	sqlReq = "update book set bookname = $1 where bookname =$2;"

	if _, err := conn.Exec(context.Background(), sqlReq,
		update.NewValue,
		update.OldValue); err != nil {
		fmt.Fprint(c.Writer, "Error\n")
	} else {
		fmt.Fprint(c.Writer, "Книга Обновлена!\n")
	}
	fmt.Fprint(c.Writer, sqlReq)

}
