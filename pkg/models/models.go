package models

type Book struct {
	BookName      string `json:"bookName"`
	YearOfPublish int    `json:"yearOfPublish"`
	AuthorID      int    `json:"authorID"`
	PublisherID   int    `json:"publisherID"`
}

type Publisher struct {
	ID   int
	Name string
}

type Author struct {
	ID   int
	Name string
}
