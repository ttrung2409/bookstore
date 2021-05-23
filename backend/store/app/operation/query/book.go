package query

type Book struct {
	GoogleBookId  string
	Title         string
	Subtitle      string
	Description   string
	Authors       []string
	AverageRating float32
	RatingsCount  int
	ThumbnailUrl  string
	PreviewUrl    string
	OnhandQty     int
	ReservedQty   int
}
