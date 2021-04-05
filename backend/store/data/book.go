package data

import (
	"time"
)

type Book struct {
	Id       EntityId
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
}

type BookRepository interface {
	repositoryBase
	GetByGoogleBookId(googleBookId string) (interface{}, error)
	CreateIfNotExist(book Book, transaction *Transaction) error
}