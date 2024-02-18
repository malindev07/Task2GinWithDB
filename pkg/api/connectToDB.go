package api

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func ConnectDB() (*pgx.Conn, error) {
	var DbUrl = "postgres://postgres:Respekt42@localhost:5432/BookShop"
	conn, err := pgx.Connect(context.Background(), DbUrl)

	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("All good")
	}

	return conn, err
}
