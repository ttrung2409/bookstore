package query

type Query interface {
	SearchGoogleBooks(term string) ([]Book, error)
}

type query struct{}
