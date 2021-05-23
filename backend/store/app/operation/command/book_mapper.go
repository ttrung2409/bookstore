package operation

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

func (Book) fromDataObject(b data.Book) Book {
	return Book{
		GoogleBookId:  b.GoogleBookId,
		Title:         b.Title,
		Subtitle:      b.Subtitle,
		Description:   b.Description,
		Authors:       b.Authors,
		AverageRating: b.AverageRating,
		RatingsCount:  b.RatingsCount,
		ThumbnailUrl:  b.ThumbnailUrl,
		PreviewUrl:    b.PreviewUrl,
		ReservedQty:   b.ReservedQty,
		OnhandQty:     b.OnhandQty,
	}
}
