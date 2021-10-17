package query

type Query interface {
	FindGoogleBooks(term string) ([]Book, error)
}

type query struct{}
