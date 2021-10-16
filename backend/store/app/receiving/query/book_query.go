package query

type BookQuery interface {
	SearchGoogleBooks(term string) ([]Book, error)
}
