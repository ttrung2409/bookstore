package operation

type FindBooks interface {
	Find(term string) ([]Book, error)
}
