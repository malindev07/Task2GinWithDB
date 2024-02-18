package models

type Book struct {
	BookName      string
	YearOfPublish int
	AuthorID      int

	PublisherID int
}

type Publisher struct {
	ID   int
	Name string
}

type Author struct {
	ID   int
	Name string
}
