package query

type GoogleBookQuery interface {
	Find(term string) ([]Book, error)
}
