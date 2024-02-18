package api

import (
	"github.com/gin-gonic/gin"

	"context"
	"fmt"
	"log"
)

func ShowAuthorById(c *gin.Context) {
	id := c.Param("id")
	fmt.Fprint(c.Writer, id+"\n")

	//создаем соединение с бд
	conn, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(context.Background())

	//var author models.Author

	var res = ""
	sqlReq := "SELECT name FROM authors WHERE authors.id =" + id

	//запрос к бд с указанным id
	err = conn.QueryRow(context.Background(), sqlReq).Scan(&res)

	fmt.Fprintf(c.Writer, res)

}
