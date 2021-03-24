package domain

import (
	"repository"
	"time"
)

type Book struct {
	id            string
	title         string
	subtitle      string
	description   string
	authors       []string
	publisher string
	publishedDate time.Time
	averageRating float32
	ratingsCount int32
	thumbnailUrl string
}

func Get(id string) Book {
	var book Book
		book = repository.BookRepository.Get(id);
		return book
}

func Create(book Book) Book {
	return Book
}



func (book *Book) 