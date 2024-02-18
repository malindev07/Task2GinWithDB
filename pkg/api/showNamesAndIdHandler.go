package api

import (
	"Task2DbWithGin/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"

	"context"
	"log"
)

func ShowNameAndIdHandler(c *gin.Context) {
	conn, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(context.Background())
	var author models.Author

	rows, err := conn.Query(context.Background(), "SELECT * FROM authors;")
	// выводи все строки, предварительно поместив из  в структуру
	for rows.Next() {
		rows.Scan(
			&author.ID,
			&author.Name)
		fmt.Fprintf(c.Writer,
			"ID автора: %v, Имя: %v\n",
			author.ID, author.Name)
	}
}
