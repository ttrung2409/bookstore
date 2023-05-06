package query

import "store/app/domain"

type Book struct {
	Id           string
	Title        string
	Subtitle     string
	Description  string
	ThumbnailUrl string
	OnhandQty    int
	ReservedQty  int
}

func (Book) fromDataObject(b domain.BookData) Book {
	return Book{
		Id:           b.Id,
		Title:        b.Title,
		Subtitle:     b.Subtitle,
		Description:  b.Description,
		ThumbnailUrl: b.ThumbnailUrl,
		ReservedQty:  b.ReservedQty,
		OnhandQty:    b.OnhandQty,
	}
}
