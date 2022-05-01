package query

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"store/app/domain/data"
	repo "store/app/repository"
	"store/container"
	"store/utils"
)

type Query interface {
	FindGoogleBooks(term string) ([]*Book, error)
	FindBooks(term string) ([]*Book, error)
}

func New() Query {
	return &query{}
}

type query struct{}

func (*query) FindGoogleBooks(term string) ([]*Book, error) {
	url, err := url.Parse("https://www.googleapis.com/books/v1/volumes")
	if err != nil {
		return nil, err
	}

	query := url.Query()
	query.Set("q", term)
	url.RawQuery = query.Encode()

	res, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var books []*Book
	err = json.Unmarshal([]byte(body), &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (*query) FindBooks(term string) ([]*Book, error) {
	queryFactory := container.Instance().Get(utils.Nameof((*repo.QueryFactory)(nil))).(repo.QueryFactory)

	records, err := queryFactory.New(&data.Book{}).
		Where("title").Contains(term).
		Or("subtitle").Contains("term").
		Or("description").Contains(term).
		Find()

	if err != nil {
		return nil, err
	}

	var books []*Book
	for _, record := range records {
		book := Book{}.fromDataObject(record.(*data.Book))
		books = append(books, book)
	}

	return books, nil
}
