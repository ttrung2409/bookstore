package event

import (
	"store/app/messaging"
	"time"
)

type BookCreated struct {
	*messaging.Message
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
