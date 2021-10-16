package query

import "store/app/domain/data"

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
