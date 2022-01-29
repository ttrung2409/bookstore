package data

import (
	"time"
)

type Book struct {
	Id            EntityId `gorm:"primaryKey"`
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

func (b *Book) GetId() EntityId {
	return b.Id
}

func (b *Book) SetId(id EntityId) {
	b.Id = id
}

func (b *Book) Clone() *Book {
	return &Book{
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
