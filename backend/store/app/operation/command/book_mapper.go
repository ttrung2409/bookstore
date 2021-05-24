package command

import (
	"store/app/data"
)

func (b Book) toDataObject() data.Book {
	return data.Book{
		GoogleBookId:  b.GoogleBookId,
		Title:         b.Title,
		Subtitle:      b.Subtitle,
		Description:   b.Description,
		Authors:       b.Authors,
		AverageRating: b.AverageRating,
		RatingsCount:  b.RatingsCount,
		ThumbnailUrl:  b.ThumbnailUrl,
		PreviewUrl:    b.PreviewUrl,
	}
}
