package data

import (
	"fmt"
	"strings"
	"time"
)

type BookId EntityId

func (id BookId) Value() EntityId {
	return (EntityId)(id)
}

func (id BookId) ToMap() map[string]interface{} {
	return map[string]interface{}{"store_id": StoreId(), "google_book_id": googleBookId(id)}
}

func googleBookId(id BookId) string {
	return strings.Split(id.Value().ToString(), "@")[1]
}

func NewBookId(googleBookId string) BookId {
	return (BookId)(fmt.Sprintf("%s@%s", StoreId(), googleBookId))
}

func EmptyBookId() BookId {
	return BookId(EmptyEntityId)
}

type Book struct {
	Id            BookId
	StoreId       EntityId
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
	PreviewUrl    string
	OnhandQty     int
	PreservedQty  int
}

type BookRepository interface {
	repositoryBase
	UpdateOnhandQty(id BookId, qty, transaction *Transaction) error
}
