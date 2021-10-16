package query

import "store/app/domain/data"

func (Book) fromDataObject(b data.Book) Book {
	return Book{
		Id:           b.Id.ToString(),
		Title:        b.Title,
		Subtitle:     b.Subtitle,
		Description:  b.Description,
		ThumbnailUrl: b.ThumbnailUrl,
		ReservedQty:  b.ReservedQty,
		OnhandQty:    b.OnhandQty,
	}
}
