package query

import (
	"store/app/domain"
	"time"
)

type Book struct {
	Id            string
	Title         string
	Subtitle      string
	Description   string
	Authors       []string
	Publisher     string
	PublishedDate time.Time
	AverageRating float32
	RatingsCount  int
	ThumbnailUrl  string
	PreviewUrl    string
}

func (Book) fromDataObject(book domain.BookData) Book {
	return Book{
		Id:            book.Id,
		Title:         book.Title,
		Subtitle:      book.Subtitle,
		Description:   book.Description,
		Authors:       book.Authors,
		Publisher:     book.Publisher,
		PublishedDate: book.PublishedDate,
		AverageRating: book.AverageRating,
		RatingsCount:  book.RatingsCount,
		ThumbnailUrl:  book.ThumbnailUrl,
		PreviewUrl:    book.PreviewUrl,
	}
}
