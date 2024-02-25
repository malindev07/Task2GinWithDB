package models

type Book struct {
	BookName      string `json:"bookName"`
	YearOfPublish int    `json:"yearOfPublish"`
	AuthorID      int    `json:"authorID"`
	PublisherID   int    `json:"publisherID"`
}

type Publisher struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OldNewValue struct {
	OldValue string `json:"oldValue"`
	NewValue string `json:"newValue"`
}
