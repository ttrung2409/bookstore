package operation

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type findBooks struct{}

func (s *findBooks) Find(term string) ([]Book, error) {
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

	var books []Book
	err = json.Unmarshal([]byte(body), &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}
