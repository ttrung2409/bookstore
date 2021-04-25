package operation

type Book struct {
	GoogleBookId  string
	Title         string
	Subtitle      string
	Description   string
	Authors       []string
	AverageRating float32
	RatingsCount  int32
	ThumbnailUrl  string
	PreviewUrl    string
	OnhandQty     int
	PreservedQty  int
}
