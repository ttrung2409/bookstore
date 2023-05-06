package domain

import (
	"time"
)

type BookData struct {
	Id            string `gorm:"primaryKey"`
	GoogleBookId  string
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
	OnhandQty     int
	ReservedQty   int
}

func (b BookData) Clone() BookData {
	return BookData{
		Id:            b.Id,
		GoogleBookId:  b.GoogleBookId,
		Title:         b.Title,
		Subtitle:      b.Subtitle,
		Description:   b.Description,
		Authors:       b.Authors,
		Publisher:     b.Publisher,
		PublishedDate: b.PublishedDate,
		AverageRating: b.AverageRating,
		RatingsCount:  b.RatingsCount,
		ThumbnailUrl:  b.ThumbnailUrl,
		PreviewUrl:    b.PreviewUrl,
		OnhandQty:     b.OnhandQty,
		ReservedQty:   b.ReservedQty,
	}
}

type Book struct {
	state BookData
}

func (Book) New(book BookData) *Book {
	cloned := book.Clone()
	if cloned.Id == "" {
		cloned.Id = NewId()
	}

	return &Book{state: cloned}
}

func (b *Book) State() BookData {
	return b.state.Clone()
}
