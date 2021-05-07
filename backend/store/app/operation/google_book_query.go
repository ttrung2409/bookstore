package operation

type GoogleBookQuery interface {
	Find(term string) ([]Book, error)
}
