package domain

import (
	"errors"
	module "store"
	repository "store/repository/interface"
	"store/utils"
	"time"
)

type Book struct {
	Id            repository.EntityId
	StoreId repository.EntityId
	GoogleBookId string
	Title         string
	Subtitle      string
	Description   string
	Authors       []string
	Publisher string
	PublishedDate time.Time
	AverageRating float32
	RatingsCount int32
	ThumbnailUrl string
}

var bookRepositoryRef = module.Container.Get(utils.Nameof((*repository.BookRepository)(nil))).(*repository.BookRepository);
var bookRepository = *bookRepositoryRef;

func GetBook(id repository.EntityId) (*Book, error) {
	book, ok := bookRepository.Get(id).(*Book)

	if (!ok) {
		return nil, errors.New("No book found")
	}

	return book, nil
}

func CreateBookIfNotExist(book Book) *Book {
	bookRepository.GetByGoogleBookId(book);
	id := bookRepository.Create(book, nil)

	createdBook := book
	createdBook.Id = id

	return &createdBook
}

func (book *Book) Update(entity Book) {
	bookRepository.Update(book.Id, book, nil);
} 