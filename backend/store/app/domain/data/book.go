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
