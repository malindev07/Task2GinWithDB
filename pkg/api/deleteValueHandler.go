package api

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"context"
	"log"
)

func DeleteValueHandler(c *gin.Context) {

	conn, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(context.Background())

	//var book models.Book

	bookName := c.Param("bookName")
	fmt.Fprint(c.Writer, bookName+"\n")

	sqlReq := "delete from book where bookname = $1;"

	if _, err := conn.Exec(context.Background(), sqlReq, bookName); err != nil {
		fmt.Fprint(c.Writer, "Error\n")
	} else {
		fmt.Fprint(c.Writer, "Книга Удалена!\n")
	}
	fmt.Fprint(c.Writer, sqlReq)
}
