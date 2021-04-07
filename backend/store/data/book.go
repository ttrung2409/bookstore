package data

import (
	"time"
)

type Book struct {
	StoreId EntityId
	GoogleBookId  string
	Title         string
	Subtitle      string
	Description   string
	Authors       []string
	Publisher     string
	PublishedDate time.Time
	AverageRating float32
	RatingsCount  int32
	ThumbnailUrl  string
	OnhandQty int
	PreservedQty int
}

type BookRepository interface {
	repositoryBase
	MakeId(googleBookId string) string
	CreateIfNotExist(book Book, transaction *Transaction) error
}