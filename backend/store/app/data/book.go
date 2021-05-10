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
	PreservedQty  int
}

func (b *Book) GetId() (EntityId, string) {
	return b.Id, "id"
}

func (b *Book) SetId(id EntityId) {
	b.Id = id
}

type BookRepository interface {
	repositoryBase
	CreateIfNotExists(book Book, tx Transaction) (EntityId, error)
	AdjustOnhandQty(id EntityId, qty int, tx Transaction) error
}
