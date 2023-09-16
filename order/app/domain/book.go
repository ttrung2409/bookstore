package domain

type BookData struct {
	Id           string
	Title        string
	Subtitle     string
	Description  string
	ThumbnailUrl string
}

func (this *BookData) Clone() BookData {
	return BookData{
		Id:           this.Id,
		Title:        this.Title,
		Subtitle:     this.Subtitle,
		Description:  this.Description,
		ThumbnailUrl: this.ThumbnailUrl,
	}
}

type Book struct {
	book BookData
}

func (Book) New(book BookData) *Book {
	return &Book{book: book}
}

func (book *Book) State() BookData {
	return book.book.Clone()
}
