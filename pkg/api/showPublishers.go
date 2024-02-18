package api

import (
	"Task2DbWithGin/pkg/models"

	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func ShowPublisher(c *gin.Context) {
	// создаем соединение с бд
	conn, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(context.Background())

	var publisher models.Publisher

	// создаем запрос к бд
	rows, err := conn.Query(context.Background(), "SELECT * FROM publisher;")
	// выводи все строки, предварительно поместив из  в структуру
	for rows.Next() {
		rows.Scan(
			&publisher.ID,
			&publisher.Name)
		fmt.Fprintf(c.Writer,
			"ID издателя: %v, Название: %v\n",
			publisher.ID, publisher.Name)
	}
}
