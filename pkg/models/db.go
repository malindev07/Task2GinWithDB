package models

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

var Connect *pgx.Conn

func SetConnect(dbUrl *string) error {

	conn, err := pgx.Connect(context.Background(), *dbUrl)
	_ = conn
	defer conn.Close(context.Background())

	Connect = conn

	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("Соединение с базой данных произошло успешно!")
	}

	return err

}

func GetConnect() *pgx.Conn {
	return Connect
}

func GetBooks(book Book) []Book {

	rows, err := Connect.Query(context.Background(), "SELECT * FROM book;")
	var books []Book
	// выводи все строки, предварительно поместив иx  в структуру
	for rows.Next() {
		rows.Scan(
			&book.BookName,
			&book.YearOfPublish,
			&book.AuthorID,
			&book.PublisherID)
		books = append(books, book)
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	return books
}

func DeleteValue(param string) string {

	//var book models.Book
	fmt.Println(param + "\n")

	sqlReq := "delete from book where bookname = $1;"

	_, err := Connect.Exec(context.Background(), sqlReq, param)
	if err != nil {
		log.Fatal(err.Error())
	}

	return "Книга с названием " + param + " удалена!"

}

func ShowAuthorById(id string) Author {

	var author Author

	sqlReq := "SELECT * FROM authors WHERE authors.id =" + id

	//запрос к бд с указанным id

	err := Connect.QueryRow(context.Background(), sqlReq).Scan(&author.ID, &author.Name)

	if err != nil {
		log.Fatal(err.Error())
	}

	return author
}

func ShowAuthors() []Author {

	var author Author
	var authors []Author

	rows, err := Connect.Query(context.Background(), "SELECT * FROM authors;")
	// выводи все строки, предварительно поместив из  в структуру
	for rows.Next() {
		rows.Scan(
			&author.ID,
			&author.Name)
		authors = append(authors, author)
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	return authors
}

func ShowPublisherById(id string) Publisher {

	var publisher Publisher

	sqlReq := "SELECT * FROM publisher WHERE publisher.id =" + id

	//запрос к бд с указанным id
	err := Connect.QueryRow(context.Background(), sqlReq).Scan(
		&publisher.ID,
		&publisher.Name)

	if err != nil {
		log.Fatal(err.Error())
	}

	return publisher
}

func ShowPublishers() []Publisher {

	var publisher Publisher
	var publishers []Publisher

	// создаем запрос к бд
	rows, err := Connect.Query(context.Background(), "SELECT * FROM publisher;")
	// выводи все строки, предварительно поместив из  в структуру
	for rows.Next() {
		rows.Scan(
			&publisher.ID,
			&publisher.Name)

		publishers = append(publishers, publisher)
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	return publishers
}

func AddBook(c *gin.Context) {

	var book Book

	// биндим занчения из post  в json формат и передаем в структуру book
	if err := c.BindJSON(&book); err != nil {
		fmt.Fprint(c.Writer, "Возникла ошибка")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	// тело запроса к бд
	sqlReq := "INSERT INTO book (bookname, yearofpublish, author_id, publisher_id) VALUES ($1,$2,$3,$4)"

	// отправка данных в бд
	if _, err := Connect.Exec(context.Background(), sqlReq,
		&book.BookName,
		&book.YearOfPublish,
		&book.AuthorID,
		&book.PublisherID); err != nil {
		fmt.Fprint(c.Writer, "Error\n")
	} else {
		fmt.Fprint(c.Writer, "Книга добавлена!\n")
	}
	_ = "\n"
	c.JSON(http.StatusCreated, book)
}

//func UpdateBook(c *gin.Context) {
//
//	var book Book
//	var update OldNewValue
//
//	if err := c.BindJSON(&update); err != nil {
//		fmt.Fprint(c.Writer, "Возникла ошибка")
//		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
//	}
//	sqlData := fmt.Sprintf("'%v'", update.OldValue)
//
//	sqlReq := "SELECT * FROM book WHERE bookname =" + sqlData
//
//	err := Connect.QueryRow(context.Background(), sqlReq).Scan(
//		&book.BookName,
//		&book.YearOfPublish,
//		&book.AuthorID,
//		&book.PublisherID)
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	fmt.Fprint(c.Writer, book, "\n")
//
//	sqlReq = "update book set bookname = $1 where bookname =$2;"
//
//	if _, err := Connect.Exec(context.Background(), sqlReq,
//		update.NewValue,
//		update.OldValue); err != nil {
//		fmt.Fprint(c.Writer, "Error\n")
//	} else {
//		fmt.Fprint(c.Writer, "Книга Обновлена!\n")
//	}
//	fmt.Fprint(c.Writer, sqlReq)
//
//}
