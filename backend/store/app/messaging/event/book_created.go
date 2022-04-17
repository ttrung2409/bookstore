package event

import (
	"time"
)

type BookCreated struct {
	Id            string
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
}

func (event BookCreated) Key() string {
	return event.Id
}

func (event BookCreated) Topic() string {
	return "book"
}
